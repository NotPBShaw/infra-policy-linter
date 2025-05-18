# Architecture

- `cmd/.../main.go` is CLI entrypoint.
- `internal/scanner` contains policy evaluation engine.
- Findings are returned as structured records for downstream pipelines.
