// @bun
// src/index.ts
import { readFileSync } from "fs";
import { homedir } from "os";
import { join } from "path";
var PROVIDER_ID = "closedrouter";
var API_BASE = "https://router.queef.in/v1";
var DEFAULT_CONTEXT_WINDOW = 128000;
var DEFAULT_OUTPUT_TOKENS = 8192;
function getAuthPath() {
  const dataHome = process.env.XDG_DATA_HOME || join(homedir(), ".local", "share");
  return join(dataHome, "opencode", "auth.json");
}
function readStoredKey() {
  try {
    const raw = readFileSync(getAuthPath(), "utf-8");
    const data = JSON.parse(raw);
    const entry = data?.[PROVIDER_ID];
    if (entry?.type === "api" && entry?.key)
      return entry.key;
  } catch {}
  return;
}
function normalizeModalities(values, fallback) {
  const allowed = new Set(["text", "audio", "image", "video", "pdf"]);
  const normalized = (values ?? fallback).filter((value) => allowed.has(value));
  return normalized.length > 0 ? normalized : fallback;
}
async function fetchModels(apiKey) {
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
    return json.data ?? [];
  } catch {
    return [];
  }
}
var ClosedRouterPlugin = async () => {
  return {
    config: async (config) => {
      config.provider ??= {};
      config.provider[PROVIDER_ID] ??= {};
      config.provider[PROVIDER_ID].npm = "@ai-sdk/openai";
      config.provider[PROVIDER_ID].api = API_BASE;
      config.provider[PROVIDER_ID].options ??= {};
      const key = readStoredKey();
      if (key) {
        config.provider[PROVIDER_ID].options.apiKey = key;
      } else {
        return;
      }
      const models = await fetchModels(key);
      if (models.length === 0)
        return;
      config.provider[PROVIDER_ID].models ??= {};
      for (const model of models) {
        const inputModalities = normalizeModalities(model.modalities?.input, ["text"]);
        const outputModalities = normalizeModalities(model.modalities?.output, ["text"]);
        const hasImageInput = inputModalities.includes("image");
        config.provider[PROVIDER_ID].models[model.id] = {
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
      }
      if (!config.small_model && (config.provider[PROVIDER_ID].models?.["title-gen"] || config.provider[PROVIDER_ID].models?.[`${PROVIDER_ID}/title-gen`])) {
        config.small_model = `${PROVIDER_ID}/title-gen`;
      }
    },
    auth: {
      provider: PROVIDER_ID,
      loader: async (getAuth) => {
        try {
          const stored = await getAuth();
          if (!stored || stored.type !== "api")
            return {};
          return { apiKey: stored.key };
        } catch {
          return {};
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
      const providerId = input.provider?.info?.id ?? input.provider?.id ?? input.model?.providerID;
      if (providerId !== PROVIDER_ID)
        return;
      const apiKey = input.provider.options?.apiKey ?? input.provider?.key ?? readStoredKey();
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
