/**
 * Copilot Force Agent Header Plugin for OpenCode
 * 
 * This is a modified version of opencode-copilot-auth with configurable X-Initiator behavior.
 * - For non-first messages (messages with assistant/tool roles): behaves like default plugin
 * - For first messages: has 1/USER_INITIATOR_RATIO chance of using "user", otherwise "agent"
 * 
 * Based on: https://github.com/sst/opencode-copilot-auth/blob/main/index.mjs
 */

import type { Plugin } from "@opencode-ai/plugin"
import type { OAuth } from "@opencode-ai/sdk"
import { appendFileSync } from "fs"

const DEBUG_ENABLED = true // Set to true to enable debug logging
const DEBUG_LOG = '/tmp/opencode-copilot-agent-header-debug.log'
const USER_INITIATOR_RATIO = 2 // 1/X chance of "user" for first messages (<=0 disables and always uses "agent")

// Responses API alternate input types (from copilot-auth 0.0.8)
const RESPONSES_API_ALTERNATE_INPUT_TYPES = [
  "file_search_call",
  "computer_call",
  "computer_call_output",
  "web_search_call",
  "function_call",
  "function_call_output",
  "image_generation_call",
  "code_interpreter_call",
  "local_shell_call",
  "local_shell_call_output",
  "mcp_list_tools",
  "mcp_approval_request",
  "mcp_approval_response",
  "mcp_call",
  "reasoning",
]

function log(message: string) {
    if (!DEBUG_ENABLED) return
    try {
        const timestamp = new Date().toISOString()
        appendFileSync(DEBUG_LOG, `${timestamp} ${message}\n`)
    } catch (error) {
        if (DEBUG_ENABLED) {
            console.error(`[PLUGIN_LOG_ERROR] Failed to write to ${DEBUG_LOG}:`, error)
        }
    }
}

const CopilotForceAgentHeader: Plugin = async ({ client }) => {
    log('[INIT] Copilot Force Agent Header plugin loaded (replaces copilot-auth)')
    log(`[INIT] OPENCODE_DISABLE_DEFAULT_PLUGINS = ${process.env.OPENCODE_DISABLE_DEFAULT_PLUGINS || 'NOT SET'}`)

    // Constants from copilot-auth
    const HEADERS = {
        "User-Agent": "GitHubCopilotChat/0.32.4",
        "Editor-Version": "vscode/1.105.1",
        "Editor-Plugin-Version": "copilot-chat/0.32.4",
        "Copilot-Integration-Id": "vscode-chat",
    }

    function normalizeDomain(url: string) {
        return url.replace(/^https?:\/\//, "").replace(/\/$/, "")
    }

    function getUrls(domain: string) {
        return {
            COPILOT_API_KEY_URL: `https://api.${domain}/copilot_internal/v2/token`,
        }
    }

    return {
        auth: {
            provider: "github-copilot",
            loader: async (getAuth, provider) => {
                log('[AUTH_LOADER] Loading auth for github-copilot')

                const authInfo = await getAuth()
                if (!authInfo || authInfo.type !== "oauth") {
                    log('[AUTH_LOADER] No OAuth found')
                    return {}
                }

                // Type assertion after runtime check
                const info = authInfo as OAuth

                // Set model costs to 0 (from copilot-auth 0.0.7)
                if (provider && provider.models) {
                    for (const model of Object.values(provider.models)) {
                        model.cost = { input: 0, output: 0, cache_read: 0, cache_write: 0 }
                    }
                }

                // Determine baseURL
                const enterpriseUrl = info.enterpriseUrl
                const baseURL = enterpriseUrl
                    ? `https://copilot-api.${normalizeDomain(enterpriseUrl)}`
                    : "https://api.githubcopilot.com"

                log(`[AUTH_LOADER] baseURL: ${baseURL}`)

                const fetchWrapper = async (input: RequestInfo | URL, init?: RequestInit) => {
                    log('[FETCH] Fetch function called!')
                    const authInfo = await getAuth()
                    if (!authInfo || authInfo.type !== "oauth") {
                        return fetch(input, init)
                    }

                    // Type assertion after runtime check
                    const currentInfo = authInfo as OAuth

                    // Token refresh logic
                    if (!currentInfo.access || currentInfo.expires < Date.now()) {
                        log('[FETCH] Token expired, refreshing...')

                        const domain = currentInfo.enterpriseUrl
                            ? normalizeDomain(currentInfo.enterpriseUrl)
                            : "github.com"
                        const urls = getUrls(domain)

                        const response = await fetch(urls.COPILOT_API_KEY_URL, {
                            headers: {
                                Accept: "application/json",
                                Authorization: `Bearer ${currentInfo.refresh}`,
                                ...HEADERS,
                            },
                        })

                        if (!response.ok) {
                            log('[FETCH] Token refresh failed')
                            throw new Error(`Token refresh failed: ${response.status}`)
                        }

                        const tokenData = await response.json()

                        const saveProviderID = currentInfo.enterpriseUrl
                            ? "github-copilot-enterprise"
                            : "github-copilot"

                        await client.auth.set({
                            path: { id: saveProviderID },
                            body: {
                                type: "oauth",
                                refresh: currentInfo.refresh,
                                access: tokenData.token,
                                expires: tokenData.expires_at * 1000,
                                ...(currentInfo.enterpriseUrl && {
                                    enterpriseUrl: currentInfo.enterpriseUrl,
                                }),
                            },
                        })

                        currentInfo.access = tokenData.token
                        log('[FETCH] Token refreshed')
                    }

                    // Determine X-Initiator based on message content
                    let isAgentCall = false
                    let isVisionRequest = false
                    try {
                        const body = typeof init?.body === "string"
                            ? JSON.parse(init.body)
                            : init?.body
                        
                        // Check Chat Completions API format (messages array)
                        if (body?.messages) {
                            // Check if this is an ongoing conversation (has assistant/tool messages)
                            isAgentCall = body.messages.some(
                              (msg: any) => msg.role && ["tool", "assistant"].includes(msg.role),
                            )
                            // Check for vision request
                            isVisionRequest = body.messages.some(
                                (msg: any) =>
                                    Array.isArray(msg.content) &&
                                    msg.content.some((part: any) => part.type === "image_url"),
                            )
                        }
                        
                        // Check Responses API format (input array) - from copilot-auth 0.0.8
                        if (!isAgentCall && body?.input && Array.isArray(body.input)) {
                            const lastInput = body.input[body.input.length - 1]
                            const isAssistant = lastInput?.role === "assistant"
                            const hasAgentType = lastInput?.type
                                ? RESPONSES_API_ALTERNATE_INPUT_TYPES.includes(lastInput.type)
                                : false
                            isAgentCall = isAssistant || hasAgentType
                            
                            // Check for vision request in Responses API format
                            if (Array.isArray(lastInput?.content)) {
                                isVisionRequest = lastInput.content.some((part: any) => part.type === "input_image")
                            }
                        }
                    } catch { }

                    // Check if x-initiator was already set by another hook (e.g., built-in's chat.headers for subagent sessions)
                    const existingInitiator = (init?.headers as Record<string, string>)?.["x-initiator"]

                    // Determine X-Initiator value
                    let initiator: string
                    if (existingInitiator) {
                        // Respect header already set by built-in chat.headers (subagent sessions)
                        initiator = existingInitiator
                        log(`[FETCH] Respecting existing x-initiator from chat.headers: ${initiator}`)
                    } else if (isAgentCall) {
                        // Non-first message: always use "agent"
                        initiator = "agent"
                        log('[FETCH] Non-first message detected, using: agent')
                    } else {
                        // First message: 1/USER_INITIATOR_RATIO chance of "user", otherwise "agent"
                        const userProbability = USER_INITIATOR_RATIO > 0 ? 1 / USER_INITIATOR_RATIO : 0
                        const randomValue = Math.random()
                        const useUser = randomValue < userProbability
                        initiator = useUser ? "user" : "agent"
                        const thresholdLog = USER_INITIATOR_RATIO > 0 ? userProbability.toFixed(4) : 'disabled'
                        log(`[FETCH] First message detected, random=${randomValue.toFixed(4)}, threshold=${thresholdLog}, using: ${initiator}`)
                    }

                    // Build headers
                    const url = typeof input === 'string' ? input : input.toString()
                    log(`[FETCH] Request to: ${url.substring(0, 60)}...`)

                    const headers: any = {
                        ...init?.headers,
                        ...HEADERS,
                        Authorization: `Bearer ${currentInfo.access}`,
                        "Openai-Intent": "conversation-edits",
                        "X-Initiator": initiator,
                    }

                    if (isVisionRequest) {
                        headers["Copilot-Vision-Request"] = "true"
                    }

                    delete headers["x-api-key"]
                    delete headers["authorization"]

                    log(`[FETCH] ✓ X-Initiator: ${initiator}`)

                    return fetch(input, {
                        ...init,
                        headers,
                    })
                }

                log('[AUTH_LOADER] Returning config with custom fetch')

                return {
                    baseURL,
                    apiKey: "",
                    fetch: fetchWrapper,
                }
            },
            methods: [],
        },
    }
}

export default CopilotForceAgentHeader
