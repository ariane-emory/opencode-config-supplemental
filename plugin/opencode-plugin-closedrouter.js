// @bun
// src/index.ts
import { readFileSync } from "fs";
import { homedir } from "os";
import { join } from "path";
var PROVIDER_ID = "closedrouter";
var API_BASE = "https://router.queef.in/v1";
var DEFAULT_CONTEXT_WINDOW = 128000;
var DEFAULT_OUTPUT_TOKENS = 8192;
var MODEL_CACHE_TTL_MS = 60000;
var authCache = new Map;
var modelCache = new Map;
function getAuthPath() {
  const dataHome = process.env.XDG_DATA_HOME || join(homedir(), ".local", "share");
  return join(dataHome, "opencode", "auth.json");
}
function readStoredKeyFallback() {
  try {
    const raw = readFileSync(getAuthPath(), "utf-8");
    const data = JSON.parse(raw);
    const entry = data?.[PROVIDER_ID];
    if (entry?.type === "api" && typeof entry.key === "string") {
      authCache.set(PROVIDER_ID, entry.key);
      return entry.key;
    }
  } catch {}
  return;
}
function normalizeModalities(values, fallback) {
  const allowed = new Set(["text", "audio", "image", "video", "pdf"]);
  const normalized = (values ?? fallback).filter((value) => allowed.has(value));
  return normalized.length > 0 ? normalized : fallback;
}
function getProviderKey(npmPackage) {
  if (npmPackage === "@ai-sdk/openai")
    return PROVIDER_ID;
  const suffix = npmPackage.replace(/^@ai-sdk\//, "").trim().toLowerCase().replace(/[^a-z0-9._-]+/g, "-").replace(/-+/g, "-").replace(/^-|-$/g, "");
  return suffix ? `${PROVIDER_ID}-${suffix}` : PROVIDER_ID;
}
function buildBaseOptions(apiKey) {
  return {
    baseURL: API_BASE,
    ...apiKey ? { apiKey } : {}
  };
}
function getCachedApiKey() {
  return authCache.get(PROVIDER_ID) || process.env.CLOSEDROUTER_API_KEY || readStoredKeyFallback();
}
async function fetchModels(apiKey) {
  const cached = modelCache.get(apiKey);
  if (cached && cached.expiresAt > Date.now()) {
    return cached.models;
  }
  try {
    const res = await fetch(`${API_BASE}/models`, {
      headers: {
        Authorization: `Bearer ${apiKey}`,
        Accept: "application/json"
      }
    });
    if (!res.ok)
      return [];
    const json = await res.json();
    const models = json.data ?? [];
    modelCache.set(apiKey, {
      expiresAt: Date.now() + MODEL_CACHE_TTL_MS,
      models
    });
    return models;
  } catch {
    return [];
  }
}
var ClosedRouterPlugin = async (_input) => {
  return {
    config: async (config) => {
      config.provider ??= {};
      config.provider[PROVIDER_ID] ??= {};
      config.provider[PROVIDER_ID].npm = "@ai-sdk/openai";
      config.provider[PROVIDER_ID].api = API_BASE;
      config.provider[PROVIDER_ID].options = {
        ...config.provider[PROVIDER_ID].options ?? {},
        ...buildBaseOptions(getCachedApiKey())
      };
      const key = getCachedApiKey();
      if (!key)
        return;
      const models = await fetchModels(key);
      if (models.length === 0)
        return;
      let smallModelId;
      for (const model of models) {
        const npmPackage = model.aisdk_provider || "@ai-sdk/openai";
        const providerKey = getProviderKey(npmPackage);
        config.provider[providerKey] ??= {};
        config.provider[providerKey].npm = npmPackage;
        config.provider[providerKey].api = API_BASE;
        config.provider[providerKey].options = {
          ...config.provider[providerKey].options ?? {},
          ...buildBaseOptions(key)
        };
        config.provider[providerKey].models ??= {};
        const inputModalities = normalizeModalities(model.modalities?.input, ["text"]);
        const outputModalities = normalizeModalities(model.modalities?.output, ["text"]);
        const hasImageInput = inputModalities.includes("image");
        config.provider[providerKey].models[model.id] = {
          name: model.id,
          attachment: model.capabilities?.attachment ?? hasImageInput,
          reasoning: model.capabilities?.reasoning ?? false,
          tool_call: model.capabilities?.tool_call ?? true,
          modalities: {
            input: inputModalities,
            output: outputModalities
          },
          ...model.context_window != null || model.max_output_tokens != null ? {
            limit: {
              context: model.context_window ?? DEFAULT_CONTEXT_WINDOW,
              output: model.max_output_tokens ?? DEFAULT_OUTPUT_TOKENS
            }
          } : {}
        };
        if (!smallModelId && model.id === "title-gen") {
          smallModelId = `${providerKey}/title-gen`;
        }
      }
      if (!config.small_model && smallModelId) {
        config.small_model = smallModelId;
      }
    },
    auth: {
      provider: PROVIDER_ID,
      loader: async (getAuth, providerInfo) => {
        try {
          const stored = await getAuth();
          if (!stored || stored.type !== "api") {
            return buildBaseOptions();
          }
          authCache.set(providerInfo.id, stored.key);
          authCache.set(PROVIDER_ID, stored.key);
          return buildBaseOptions(stored.key);
        } catch {
          return buildBaseOptions();
        }
      },
      methods: [
        {
          type: "api",
          label: "API Key",
          prompts: [
            {
              type: "text",
              key: "apiKey",
              message: "ClosedRouter API key",
              placeholder: "cr-...",
              validate: (value) => value.trim().length === 0 ? "API key must not be empty" : undefined
            }
          ],
          async authorize(inputs) {
            const apiKey = inputs?.apiKey?.trim();
            if (!apiKey)
              return { type: "failed" };
            return { type: "success", key: apiKey };
          }
        }
      ]
    },
    "chat.headers": async (input, output) => {
      const providerId = input.provider.info?.id ?? input.model?.providerID;
      if (typeof providerId !== "string" || !providerId.startsWith(PROVIDER_ID))
        return;
      const apiKey = (typeof input.provider.options?.apiKey === "string" ? input.provider.options.apiKey : undefined) ?? authCache.get(providerId) ?? authCache.get(PROVIDER_ID) ?? readStoredKeyFallback();
      if (!apiKey)
        return;
      output.headers.Authorization = `Bearer ${apiKey}`;
    }
  };
};
var src_default = ClosedRouterPlugin;
export {
  src_default as default,
  ClosedRouterPlugin
};
