# reporter

A Go CLI to capture daily summaries and generate weekly summaries, grouped by project, using a locally running Ollama LLM.

## Features
- `reporter add` save a summary for today or a specific date (from flag, stdin, or $EDITOR)
- `reporter summarize` aggregate a given week (or current week) and ask Ollama to summarize
- `reporter list` lists reports for a given day (default today)
- `reporter help` for a usage instructions.
- CLI logic handled by [CobraCLI](https://github.com/spf13/cobra)
- [SQLite](https://github.com/glebarez/sqlite) storage interfaced with [GORM](https://github.com/go-gorm/gorm) (DB file: `~/.config/reporter/data.db`)
- Configurable LLM models

## Installation
Installation is available through `go install`:
```shell
go install github.com/eric-sims/reporter@latest
```

## Connect Ollama
[Instructions](https://github.com/ollama/ollama/blob/main/README.md#quickstart) for installing Ollama, finding models and serving it locally.

Set these environment variables (or flags):

`OLLAMA_HOST  (flag -ollama)` \
`OLLAMA_MODEL (flag -model)`

## Upcoming features
- [ ] Ability to connect to ChatGPT API
