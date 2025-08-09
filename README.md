# reporter — daily summaries CLI (projects & tags)

A Go CLI to capture daily summaries and generate weekly summaries, grouped by project, using a locally running Ollama LLM.

## Features
- `reporter add` save a summary for today or a specific date (from flag, stdin, or $EDITOR)
  *NEW*: add one or more --project and --tag values
- `reporter week` aggregate a given week (or current week) and ask Ollama to summarize
  *NEW*: `--per-project` groups notes by project in the prompt (default true)
- SQLite storage via GORM (DB file: ~/.config/reporter/data.db)
- Configurable Ollama host and model
