package topic08_packages_modules

/*
Title: EX08 — Workspace and multi‑module guidance (text)

Why this matters
- Senior engineers must understand how to structure multi‑module repositories and use workspaces without breaking builds.

Requirements
- Implement WorkspaceGuidance() string returning a multi‑paragraph guidance that covers at least:
  - `go work` overview and when to use it
  - `go mod init` / `go mod tidy` in submodules
  - Using `replace` directives only for local development
  - Avoiding import cycles across modules
  - Running `go list` / `go test ./...` inside the workspace

Constraints and pitfalls
- Keep the guidance under ~500 words, but mention all required bullets verbatim or closely.
- Plain text only.

Tricky edge case
- The text should mention how to run a package in a submodule without the workspace (standalone) and with workspace enabled.
*/

func WorkspaceGuidance() string { // TODO: implement guidance text
	return ""
}
