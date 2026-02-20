/**
 * Audio notification plugin
 *
 * Plays audio notifications for OpenCode events:
 * 1. When the agent finishes working (busy→idle transition)
 * 2. When a permission prompt is displayed to the user
 *
 * How it works:
 * - Tracks session status transitions (idle, busy, retry)
 * - Rings the bell when transitioning from "busy" to "idle"
 * - Plays permission sound when permission.updated event fires
 * - This indicates the agent has completed processing
 *
 * Why this approach works:
 * - The OpenCode session loop sets status to "busy" at the start of processing
 * - When processing completes (success, error, or cancellation), status is set to "idle"
 * - By detecting the busy→idle transition, we reliably catch completion events
 * - The permission.updated event fires right before the modal is shown to the user
 * - This avoids false positives from other status changes
 */
import type { Plugin } from "@opencode-ai/plugin";
import { spawn } from "child_process";
import { accessSync, constants, writeFileSync, unlinkSync, existsSync } from "fs";
import { resolve, normalize } from "path";

const configPath = getConfigPath();
const semaphoreFile = normalize(resolve(configPath, "prompt-enhancer", "prompt-enhancer-semaphore"),);
const semaphorePresentSound = [normalize(resolve(configPath, "audio", "riff.aiff")), 0.3,];
const noSemaphorePresentSound = [normalize(resolve(configPath, "audio", "riff2.aiff")), 0.3,];
const permissionPromptSound = [normalize(resolve(configPath, "audio", "riff3.aiff")), 0.3,];
const questionAskedSound = permissionPromptSound; // Same sound for now, can be customized later

// Audio queue and coordination system
interface SoundRequest {
  filePath: string;
  volume: number;
}

let audioQueue: SoundRequest[] = [];
let isProcessingQueue = false;
let currentProcess: any = null;

// Check if a session belongs to a primary agent (no parentID) vs subagent (has parentID)
async function isPrimaryAgentSession(sessionID: string, client: any): Promise<boolean> {
  const debugMode = process.env.OPENCODE_DEBUG_BELL === 'true';
  
  try {
    // Query session details using the client
    // SDK v1: use { path: { id: sessionID } } and returns { data: session }
    const response = await client.session.get({ path: { id: sessionID } });
    const session = response.data;
    
    if (debugMode) {
      process.stderr.write(`TerminalBell: [DEBUG] Session ${sessionID} response: ${JSON.stringify(session, null, 2)}\n`);
    }
    
    // Check if session was returned properly
    if (!session || typeof session !== 'object') {
      if (debugMode) {
        process.stderr.write(`TerminalBell: [DEBUG] Session ${sessionID} returned invalid data, assuming sub-session (no sound)\n`);
      }
      return false;
    }
    
    // Primary agents have no parentID, subagents do
    const isPrimary = !session.parentID;
    
    // Debug logging to help troubleshoot (only when enabled)
    if (debugMode) {
      process.stderr.write(`TerminalBell: [DEBUG] Session ${sessionID} parentID: ${session.parentID}, isPrimary: ${isPrimary}\n`);
    }
    return isPrimary;
  } catch (error) {
    // If we can't determine, default to false (no sound) to prevent sub-session sounds
    // This is safer than playing sounds for unknown sessions
    if (debugMode) {
      process.stderr.write(`TerminalBell: [ERROR] Failed to get session ${sessionID}: ${error instanceof Error ? error.message : String(error)}\n`);
    }
    return false;
  }
}

// Enhanced path resolution with validation
function getConfigPath(): string {
  const home = process.env.HOME || process.env.USERPROFILE || "";
  if (!home) {
    throw new Error("Home directory not found in environment variables");
  }
  return normalize(resolve(home, ".config", "opencode"));
}

function validateFileAccess(filePath: string): {
  accessible: boolean;
  error?: string;
} {
  try {
    // Check if file exists and is readable
    accessSync(filePath, constants.R_OK);
    return { accessible: true };
  } catch (error) {
    return {
      accessible: false,
      error: error instanceof Error ? error.message : String(error),
    };
  }
}


// Lock file path for cross-instance coordination
const audioLockFile = normalize(resolve(getConfigPath(), ".audio-lock"));

function acquireAudioLock(): boolean {
  try {
    // Try to create lock file with current PID
    writeFileSync(audioLockFile, process.pid.toString(), { flag: 'wx' });
    return true;
  } catch (error) {
    // Lock file exists, check if the process is still running
    try {
      if (existsSync(audioLockFile)) {
        const lockContent = require('fs').readFileSync(audioLockFile, 'utf8');
        const lockPid = parseInt(lockContent.trim(), 10);
        
        // Check if process with this PID exists (Unix systems)
        try {
          process.kill(lockPid, 0);
          // Process still exists, can't acquire lock
          return false;
        } catch (killError) {
          // Process doesn't exist, we can take over the lock
          try {
            writeFileSync(audioLockFile, process.pid.toString());
            return true;
          } catch (writeError) {
            return false;
          }
        }
      }
    } catch (readError) {
      // Can't read lock file, assume it's locked
      return false;
    }
    return false;
  }
}

function releaseAudioLock(): void {
  try {
    if (existsSync(audioLockFile)) {
      const lockContent = require('fs').readFileSync(audioLockFile, 'utf8');
      const lockPid = parseInt(lockContent.trim(), 10);
      
      // Only remove lock if we own it
      if (lockPid === process.pid) {
        unlinkSync(audioLockFile);
      }
    }
  } catch (error) {
    // Ignore errors when releasing lock
  }
}

function isAfplayRunning(): boolean {
  try {
    const result = require('child_process').spawnSync('pgrep', ['-x', 'afplay'], { 
      encoding: 'utf8',
      timeout: 1000 
    });
    return result.status === 0;
  } catch {
    return false;
  }
}

async function waitForLockAndPlay(filePath: string, volume: number): Promise<void> {
  const debugMode = process.env.OPENCODE_DEBUG_BELL === 'true';
  
  return new Promise((resolve) => {
    const tryPlay = () => {
      if (isAfplayRunning()) {
        if (debugMode) {
          process.stderr.write(`TerminalBell: [DEBUG] afplay already running, waiting...\n`);
        }
        setTimeout(tryPlay, 50);
        return;
      }
      
      if (debugMode) {
        process.stderr.write(`TerminalBell: [DEBUG] Playing ${filePath}\n`);
      }
      playSoundInternal(filePath, volume, resolve);
    };
    
    tryPlay();
  });
}

function playSoundInternal(filePath: string, volume: number, onComplete: () => void): void {
  // Validate file accessibility before attempting to play
  const fileValidation = validateFileAccess(filePath);
  if (!fileValidation.accessible) {
    // Silently fail - file access issues are typically configuration problems
    // that should be fixed at setup time, not runtime
    onComplete();
    return;
  }

  const args = [filePath, "-v", volume.toString()];

  // Spawn afplay with explicit working directory independence
  const child = spawn("afplay", args, {
    stdio: "ignore",
    detached: true,
    // Explicitly set cwd to prevent working directory dependencies
    cwd: "/",
    // Ensure clean environment
    env: { ...process.env, PATH: process.env.PATH },
  });

  // Store reference to current process
  currentProcess = child;

  // Handle critical spawn errors only
  child.on("error", (err) => {
    // Write critical errors to stderr for debug console
    process.stderr.write(
      `TerminalBell: [ERROR] Failed to spawn afplay: ${err.message}\n`,
    );
    // Clear reference on error and call onComplete
    if (currentProcess === child) {
      currentProcess = null;
    }
    onComplete();
  });

  // Clear reference when process exits
  child.on("exit", () => {
    if (currentProcess === child) {
      currentProcess = null;
    }
    onComplete();
  });

  // Unref so the parent process can exit without waiting
  child.unref();
}

function playSound(filePath: string, volume: number): void {
  // Add to queue
  audioQueue.push({ filePath, volume });
  
  if (!isProcessingQueue) {
    isProcessingQueue = true;
    processQueue();
  }
}

async function processQueue(): Promise<void> {
  const debugMode = process.env.OPENCODE_DEBUG_BELL === 'true';
  
  while (audioQueue.length > 0) {
    const request = audioQueue.shift()!;
    
    if (debugMode) {
      process.stderr.write(`TerminalBell: [DEBUG] Processing queued sound: ${request.filePath}\n`);
    }
    
    // Wait for lock and play sound
    await waitForLockAndPlay(request.filePath, request.volume);
  }
  
  isProcessingQueue = false;
}

function isSemaphorePresent(): boolean {
  // Check if semaphore file exists and is accessible
  // Missing semaphore file is expected behavior (not an error)
  const validation = validateFileAccess(semaphoreFile);
  return validation.accessible;
}

export const TerminalBell: Plugin = async ({
  project,
  client,
  $,
  directory,
  worktree,
}) => {
  // Early exit on non-macOS platforms (afplay is macOS-only)
  if (process.platform !== 'darwin') {
    return {};
  }

  // Track the last status for each session to detect busy→idle transitions
  const sessionStatus = new Map<string, "idle" | "busy" | "retry">();

  // Ensure lock is released when plugin shuts down
  const cleanup = () => {
    releaseAudioLock();
  };

  // Register cleanup handlers
  process.on('exit', cleanup);
  process.on('SIGINT', cleanup);
  process.on('SIGTERM', cleanup);
  process.on('SIGQUIT', cleanup);

  return {
    event: async ({ event }) => {
      if (event.type === "session.status") {
        const sessionId = event.properties.sessionID;
        const newStatus = event.properties.status.type;
        const previousStatus = sessionStatus.get(sessionId);

        // Update the status tracking
        sessionStatus.set(sessionId, newStatus);

        // Ring the bell when transitioning from busy to idle
        // This indicates the agent has finished working
        if (newStatus === "idle" && previousStatus === "busy") {
          // Only play sound for primary agents, not subagents
          const isPrimary = await isPrimaryAgentSession(sessionId, client);
          if (isPrimary) {
            const [soundFile, soundVolume] = isSemaphorePresent()
              ? semaphorePresentSound
              : noSemaphorePresentSound;

            playSound(soundFile, soundVolume);
          }

          // Clean up old session tracking after a delay to prevent memory leaks
          // We don't remove immediately because the session might continue
          setTimeout(() => {
            // Only delete if still idle (session hasn't restarted)
            if (sessionStatus.get(sessionId) === "idle") {
              sessionStatus.delete(sessionId);
            }
          }, 300000); // 5 minutes cleanup delay
        }
      }

      // Play sound when permission prompt is displayed
      // This alerts the user that their input is required
      if (event.type === "permission.updated") {
        // Only play sound for primary agents, not subagents
        const isPrimary = await isPrimaryAgentSession(event.properties.sessionID, client);
        if (isPrimary) {
          const [soundFile, soundVolume] = permissionPromptSound;
          playSound(soundFile, soundVolume);
        }
      }

      // Play sound when a question is asked by a tool (e.g., plan_exit)
      // This alerts the user that their input is required for mode switching
      if (event.type === "question.asked") {
        // Only play sound for questions from tool calls, not regular questions
        if (event.properties.tool) {
          // Only play sound for primary agents, not subagents
          const isPrimary = await isPrimaryAgentSession(event.properties.sessionID, client);
          if (isPrimary) {
            const [soundFile, soundVolume] = questionAskedSound;
            playSound(soundFile, soundVolume);
          }
        }
      }
    },
  };
};
