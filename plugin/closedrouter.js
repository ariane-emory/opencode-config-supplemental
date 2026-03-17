// @bun
// src/index.ts
import { readFileSync } from "fs";
import { homedir } from "os";
import { join } from "path";
var PROVIDER_ID = "closedrouter";
var API_BASE = "https://router.queef.in/v1";
function readStoredKey() {
  try {
    const authPath = join(homedir(), ".local", "share", "opencode", "auth.json");
    const raw = readFileSync(authPath, "utf-8");
    const data = JSON.parse(raw);
    const entry = data?.[PROVIDER_ID];
    if (entry?.type === "api" && entry?.key)
      return entry.key;
  } catch {}
  return;
}
async function fetchModels(apiKey) {
  const res = await fetch(`${API_BASE}/models`, {
    headers: { Authorization: `Bearer ${apiKey}` }
  });
  if (!res.ok)
    return [];
  const json = await res.json();
  return json.data ?? [];
}
var ClosedRouterPlugin = async () => {
  return {
    config: async (config) => {
      const key = readStoredKey();
      if (!key)
        return;
      config.provider ??= {};
      config.provider[PROVIDER_ID] ??= {};
      config.provider[PROVIDER_ID].api = API_BASE;
      config.provider[PROVIDER_ID].options ??= {};
      config.provider[PROVIDER_ID].options.apiKey = key;
      const models = await fetchModels(key);
      if (models.length === 0)
        return;
      config.provider[PROVIDER_ID].models ??= {};
      for (const m of models) {
        config.provider[PROVIDER_ID].models[m.id] = {
          name: m.id,
          attachment: m.capabilities?.attachment ?? m.modalities?.input?.includes("image") ?? false,
          reasoning: m.capabilities?.reasoning ?? false,
          tool_call: m.capabilities?.tool_call ?? true,
          ...m.context_window || m.max_output_tokens ? {
            limit: {
              context: m.context_window ?? 128000,
              output: m.max_output_tokens ?? 8192
            }
          } : {}
        };
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
              message: "ClosedRouter API key (from `bun run bootstrap`)",
              placeholder: "cr-...",
              validate: (v) => v.trim().length === 0 ? "API key must not be empty" : undefined
            }
          ],
          async authorize(inputs) {
            return { type: "success", key: inputs.apiKey.trim() };
          }
        }
      ]
    },
    "chat.headers": async (input, output) => {
      const providerId = input.provider?.info?.id ?? input.provider?.id ?? input.model?.providerID;
      if (providerId !== PROVIDER_ID)
        return;
      const apiKey = input.provider.options?.["apiKey"] ?? input.provider?.key ?? readStoredKey();
      if (!apiKey)
        return;
      output.headers["Authorization"] = `Bearer ${apiKey}`;
    }
  };
};
export {
  ClosedRouterPlugin
};
