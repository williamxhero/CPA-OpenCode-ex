## Agent skills

### Issue tracker

Issues and specs live in this repository's GitHub Issues. See `docs/agents/issue-tracker.md`.

### Triage labels

Use `needs-triage`, `needs-info`, `ready-for-agent`, `ready-for-human`, and `wontfix`. See `docs/agents/triage-labels.md`.

### Domain docs

This is a single-context repository. See `docs/agents/domain.md`.
## Version

- `VERSION` holds the composed identity `{source}+{ex}`, for example `8.0.76+1.0.0`: the newest tag from massiveits/opencode-go-cliproxyapi without its `v`, plus this fork's own semantic version.
- Every requirement that changes this repository ships a version change: run `D:\WILL\AGENT\CPA\ops\update-versions.ps1 -BumpChanged` before committing, or `-BumpEx minor|major` / `-ExVersion x.y.z` when a patch is too small. Merging the source repository refreshes the source part with a plain run. The script rewrites `VERSION` and the generated Go constant that ordinary `go build` compiles.
- `D:\WILL\AGENT\CPA\ops\update-versions.ps1 -Check` must pass before the task is reported finished.
- Never edit the `version_generated.go` file by hand. A test rejects an identity that is not `{source}+{ex}`.
- `-ldflags -X` still overrides the compiled version, but local deployments must rely on the checked-in identity so the running stack reports what was merged.
