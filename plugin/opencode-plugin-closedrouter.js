// @bun
// src/types.ts
var PROVIDER_ID = "closedrouter";
var API_BASE = "https://router.queef.in/v1";

// src/models.ts
var DEFAULT_NPM = "@ai-sdk/openai";
var ALLOWED_MODALITIES = new Set(["text", "audio", "image", "video", "pdf"]);
function normalizeModalities(values, fallback) {
  const result = (values ?? fallback).filter((v) => ALLOWED_MODALITIES.has(v));
  return result.length > 0 ? result : fallback;
}
function getInterleavedReasoningField(model) {
  if (model.interleaved && typeof model.interleaved === "object" && !Array.isArray(model.interleaved)) {
    const field = model.interleaved.field;
    return field === "reasoning_content" || field === "reasoning_details" ? field : undefined;
  }
  return model.capabilities?.interleaved_reasoning ? "reasoning_content" : undefined;
}
function buildReasoningVariants(efforts) {
  if (!efforts?.length)
    return;
  return Object.fromEntries(efforts.map((effort) => [effort, { reasoningEffort: effort }]));
}
function build(model) {
  const npmPackage = model.aisdk_provider || DEFAULT_NPM;
  const inputModalities = normalizeModalities(model.modalities?.input, ["text"]);
  const outputModalities = normalizeModalities(model.modalities?.output, ["text"]);
  const hasImageInput = inputModalities.includes("image");
  const supportsReasoning = model.capabilities?.reasoning ?? false;
  const interleavedField = supportsReasoning ? getInterleavedReasoningField(model) : undefined;
  const reasoningVariants = buildReasoningVariants(model.capabilities?.reasoning_efforts);
  const needsProviderOverride = npmPackage !== DEFAULT_NPM;
  const sdkModel = {
    id: model.id,
    providerID: PROVIDER_ID,
    api: {
      id: model.id,
      url: API_BASE,
      npm: npmPackage
    },
    name: model.id,
    status: "active",
    capabilities: {
      temperature: true,
      reasoning: supportsReasoning,
      attachment: model.capabilities?.attachment ?? hasImageInput,
      toolcall: model.capabilities?.tool_call ?? true,
      input: {
        text: inputModalities.includes("text"),
        audio: inputModalities.includes("audio"),
        image: inputModalities.includes("image"),
        video: inputModalities.includes("video"),
        pdf: inputModalities.includes("pdf")
      },
      output: {
        text: outputModalities.includes("text"),
        audio: outputModalities.includes("audio"),
        image: outputModalities.includes("image"),
        video: outputModalities.includes("video"),
        pdf: outputModalities.includes("pdf")
      },
      interleaved: interleavedField ? { field: interleavedField } : false
    },
    cost: {
      input: 0,
      output: 0,
      cache: {
        read: 0,
        write: 0
      }
    },
    limit: {
      context: model.context_window,
      input: undefined,
      output: model.max_output_tokens
    },
    options: {
      ...reasoningVariants ? { variants: reasoningVariants } : {}
    },
    headers: {},
    release_date: "",
    ...needsProviderOverride ? { provider: { npm: npmPackage } } : {}
  };
  if (reasoningVariants) {
    sdkModel.variants = reasoningVariants;
  }
  return sdkModel;
}
async function get(baseURL, headers = {}, existing = {}) {
  const res = await fetch(`${baseURL}/models`, {
    headers,
    signal: AbortSignal.timeout(5000)
  });
  if (!res.ok) {
    throw new Error(`Failed to fetch models: ${res.status}`);
  }
  const data = await res.json();
  const remoteModels = data.data ?? [];
  const result = { ...existing };
  for (const model of remoteModels) {
    if (model.id && model.context_window != null && model.max_output_tokens != null) {
      result[model.id] = build(model);
    }
  }
  return result;
}

// src/index.ts
import { readFileSync } from "fs";
import { appendFileSync } from "fs";
var LOG_FILE = "/tmp/closedrouter-debug.log";
function log(message, level = "DEBUG") {
  const line = `[${new Date().toISOString()}] [closedrouter-plugin ${level}] ${message}
`;
  try {
    appendFileSync(LOG_FILE, line);
  } catch {}
}
function getAuthKey() {
  try {
    const authPath = `${process.env.HOME}/.local/share/opencode/auth.json`;
    const data = JSON.parse(readFileSync(authPath, "utf-8"));
    const auth = data[PROVIDER_ID];
    if (auth?.type === "api" && auth.key) {
      return auth.key;
    }
    return;
  } catch {
    return;
  }
}
async function ClosedRouterPlugin(input, _options) {
  let models = {};
  log("=== Plugin starting ===");
  return {
    config: async (config) => {
      log("config hook called");
      config.provider ??= {};
      config.provider[PROVIDER_ID] ??= {};
      config.provider[PROVIDER_ID].npm = "@ai-sdk/openai";
      config.provider[PROVIDER_ID].api = API_BASE;
      log(`config set: provider.${PROVIDER_ID}.api = ${API_BASE}`);
      const apiKey = getAuthKey();
      if (!apiKey) {
        log("no auth key found in auth.json");
        return;
      }
      log(`fetching models with auth key (prefix: ${apiKey.substring(0, 10)}...)`);
      try {
        const result = await get(API_BASE, {
          Authorization: `Bearer ${apiKey}`,
          Accept: "application/json"
        });
        log(`fetched ${Object.keys(result).length} models`);
        for (const [id, model] of Object.entries(result)) {
          log(`  model: ${id} -> api.id=${model.api.id} npm=${model.api.npm}`);
        }
        config.provider[PROVIDER_ID].models ??= {};
        for (const [id, model] of Object.entries(result)) {
          config.provider[PROVIDER_ID].models[id] = model;
        }
        log(`added ${Object.keys(result).length} models to provider config`);
      } catch (error) {
        log(`failed to fetch models: ${error instanceof Error ? error.message : String(error)}`, "ERROR");
      }
    },
    auth: {
      provider: PROVIDER_ID,
      async loader(getAuth) {
        const info = await getAuth();
        log(`auth loader called, auth.type=${info?.type || "none"}`);
        if (!info || info.type !== "api")
          return {};
        log(`returning auth with key prefix: ${info.key.substring(0, 10)}...`);
        return {
          apiKey: "",
          async fetch(request, init) {
            const info2 = await getAuth();
            if (info2.type !== "api")
              return fetch(request, init);
            const headers = {
              ...init?.headers,
              Authorization: `Bearer ${info2.key}`
            };
            delete headers["x-api-key"];
            delete headers["authorization"];
            return fetch(request, {
              ...init,
              headers
            });
          }
        };
      },
      methods: [
        {
          type: "api",
          label: "API Key"
        }
      ]
    }
  };
}
var src_default = ClosedRouterPlugin;
export {
  src_default as default,
  ClosedRouterPlugin
};
