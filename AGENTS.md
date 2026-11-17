## Agent skills

### Issue tracker

Issues and specs live in this repository's GitHub Issues. See `docs/agents/issue-tracker.md`.

### Triage labels

Use `needs-triage`, `needs-info`, `ready-for-agent`, `ready-for-human`, and `wontfix`. See `docs/agents/triage-labels.md`.

### Domain docs

This is a single-context repository. See `docs/agents/domain.md`.
## Version

- `VERSION` holds the composed identity `{source}+{ex}`, for example `8.0.76+1.0.0`: the newest tag from massiveits/opencode-go-cliproxyapi without its `v`, plus this fork's own semantic version.
- After merging the source repository, run `D:\WILL\AGENT\CPA\ops\update-ex-versions.ps1` from the workspace root. It rewrites `VERSION` and the generated Go constant that ordinary `go build` compiles; use `-BumpEx patch` (or `-ExVersion x.y.z`) when this fork ships its own change.
- Never edit the `version_generated.go` file by hand. A test rejects an identity that is not `{source}+{ex}`.
- `-ldflags -X` still overrides the compiled version, but local deployments must rely on the checked-in identity so the running stack reports what was merged.
