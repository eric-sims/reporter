# reporter

A CLI tool to capture daily summaries and generate weekly summaries, using a locally running Ollama LLM. Written in Golang.

## Features
- Save a summary for today or a specific date
- Aggregate a given week (or current week) and ask a large language model to summarize
- List reports for a given day or week
- `reporter help` for a usage instructions.
- CLI logic handled by [CobraCLI](https://github.com/spf13/cobra)
- [SQLite](https://github.com/glebarez/sqlite) storage interfaced with [GORM](https://github.com/go-gorm/gorm) (DB file: `~/.config/reporter/data.db`)
- Configurable LLM models for Ollama, ability to access OpenAI API

## Installation
Installation is available through `go install`:
```shell
go install github.com/eric-sims/reporter@latest
```

## Connect to Ollama
[Instructions](https://github.com/ollama/ollama/blob/main/README.md#quickstart) for installing Ollama, finding models and serving it locally.

Set these environment variables (or flags):

`OLLAMA_HOST  or flag --ollama` \
`OLLAMA_MODEL or flag --model`

## Connect to OpenAI API
> [!NOTE]  
> This is different than the consumer ChatGPT subscription. This is a separate dev account separate pricing structures. Learn more about it [here](https://platform.openai.com/docs/overview).

Generate an api key [here](https://platform.openai.com/settings/organization/api-keys), then save it to the environment variable: 

`OPENAI_API_KEY or flag --openai-api-key`

## Upcoming features
- [x] Connect to ChatGPT API (user-provided api key)
- [ ] Tag projects to entries
- [ ] Easily add git-hooks to repositories to log commits
