# Modifications — AI / Cloud / Telemetry Removal

This repository is a modified fork of **Wave Terminal** (https://github.com/wavetermdev/waveterm),
originally © Command Line Inc., licensed under the **Apache License, Version 2.0**.

This file is provided to satisfy **Apache-2.0 §4(b)** ("You must cause any modified files
to carry prominent notices stating that You changed the files") and to document, for license
compliance, every category of change made in this fork. The upstream `LICENSE` and `NOTICE`
files are retained unchanged.

**Summary of intent:** remove Wave's AI features, the Wave cloud / hosted‑provider
integrations, telemetry/analytics, the network auto‑updater, and cloud marketing — leaving a
clean, offline, local‑only terminal. No code in this build contacts any Wave‑hosted service,
AI provider, telemetry endpoint, or update server.

Date of modification: 2026-06-02.

---

## 1. AI features (fully removed)

Wave's AI chat assistant, hosted/BYOK LLM provider integrations (OpenAI, Anthropic, Google
Gemini, Groq, OpenRouter, Azure, etc.), AI tool‑use, and the AI side panel were removed.

Deleted:
- `pkg/aiusechat/` — entire AI backend (provider clients, chat store, tools, prompts).
- `aiprompts/` — AI design/prompt documentation.
- `frontend/app/aipanel/` — the AI side‑panel UI and model.
- `frontend/app/view/waveai/`, `frontend/app/view/aifilediff/` — AI block views.
- `frontend/app/view/waveconfig/waveaivisual.tsx` — AI‑modes visual editor.
- `cmd/testai/`, `cmd/testopenai/`, `cmd/testsummarize/`, `cmd/wsh/cmd/wshcmd-ai.go` — AI test/CLI tools.
- `schema/waveai.json`, `schema/aipresets.json`, `pkg/wconfig/defaultconfig/waveai.json`,
  `pkg/wconfig/defaultconfig/presets/ai.json` — AI config schemas/defaults.
- AI preview harnesses under `frontend/preview/previews/`.

Surgically edited (AI wiring removed): the RPC interface (`pkg/wshrpc/wshrpctypes.go`,
`wshserver.go`), HTTP routes (`pkg/web/web.go`), config types/keys
(`pkg/wconfig/settingsconfig.go`, `metaconsts.go`), event types (`pkg/wps/wpstypes.go`),
type generators (`pkg/tsgen/*`, `cmd/generate*`), the block registry, key bindings, focus
manager, tab bars (AI button), workspace layout, Monaco schema endpoints, the term context
menu ("Send to Wave AI"), and the Tsunami builder.

The AI side‑panel layout plumbing in `workspace-layout-model.ts` is retained but **inert**:
the panel is never rendered and has no entry point (button/keybinding/RPC all removed).

## 2. Wave cloud + hosted providers (fully removed)

- Deleted `pkg/wcloud/` — all communication with Wave‑hosted endpoints
  (`api.waveterm.dev`, `ping.waveterm.dev`) including telemetry upload, diagnostic ping,
  and "no‑telemetry" notifications.
- Removed the cloud AI proxy endpoint (`cfapi.waveterm.dev`, deleted with `pkg/aiusechat`).
- Removed `wcore.GoSendNoTelemetryUpdate` and the server's cloud env‑var caching.

## 3. Telemetry / analytics (neutralized)

No telemetry is collected, stored, or transmitted.

- The `pkg/telemetry` package is reduced to **inert no‑ops** (`RecordTEvent`, `UpdateActivity`,
  etc. do nothing; `IsTelemetryEnabled()` returns `false`). Function signatures are retained so
  the many internal call sites keep compiling, but no data is recorded or persisted.
- The frontend `recordTEvent` (in `frontend/app/store/global.ts`) is an inert no‑op.
- Removed telemetry RPCs (`RecordTEventCommand`, `SendTelemetryCommand`) and the
  `telemetry:enabled` config key / opt‑in flow.
- Panic telemetry hook removed from `pkg/panichandler`.
- Note: the unused `db_activity` / `db_tevent` SQLite migrations are left in place (they only
  create empty, never‑written tables) to avoid renumbering the migration sequence.

## 4. Auto‑updater (removed)

The network auto‑updater that contacted `dl.waveterm.dev` was removed.

- Deleted `emain/updater.ts` and `frontend/app/tab/updatebanner.tsx`.
- Removed `configureAutoUpdater`, update checks, the "Check for Updates" menu item, and the
  update‑status banner from both tab bars.
- Removed the `publish` feed from `electron-builder.config.cjs`.
- The renderer's updater API surface (`getUpdaterStatus`, etc. in `emain/preload.ts`) is kept
  as inert stubs (always reports "up‑to‑date") so dependent UI compiles without network calls.
- The `autoupdate:*` settings are retained as dormant config (the `wsh` version command still
  reports a static channel label for display).

## 5. First‑run onboarding (removed)

- Deleted `frontend/app/onboarding/` and its modal registrations.
- On first launch the app now silently bootstraps the starter layout (previously triggered by
  the ToS‑agreement step in the wizard) with no wizard, telemetry opt‑in, or marketing pages.

## 6. Cloud marketing / external links (removed)

All references to `waveterm.dev` were removed from the application and Electron code:
- About modal: removed the website link and update‑channel line; de‑"AI" tagline; kept
  Open‑Source (ACKNOWLEDGEMENTS) and License links for attribution/compliance.
- Quick Tips and the durable‑session flyover: removed `docs.waveterm.dev` links.
- Help view: default URL changed from `docs.waveterm.dev` to `about:blank`.
- ARM64 translation warning: removed the docs redirect.
- `package.json`: removed `homepage`, de‑"AI‑Native" the product description.

## Build / verification notes

- The Go backend compiles cleanly (`go build ./...`).
- Generated bindings were regenerated (`task generate` / `cmd/generate*`).
- The frontend TypeScript type‑check was **not** run in this environment: dependency install is
  blocked by the repo's declared package manager (`packageManager: npm@10.9.2`), which `pnpm`
  refuses to override, and that guard was intentionally not bypassed. Frontend changes were
  verified by static reference scans instead.
