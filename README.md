# reporter — daily summaries CLI (projects & tags)
-------------------------------------------------
A Go CLI to capture daily summaries and generate weekly summaries, grouped by project, using a locally running Ollama LLM.

## Features
- `reporter add` save a summary for today or a specific date (from flag, stdin, or $EDITOR)
  *NEW*: add one or more --project and --tag values
- `reporter week` aggregate a given week (or current week) and ask Ollama to summarize
  *NEW*: `--per-project` groups notes by project in the prompt (default true)
- SQLite storage via GORM (DB file: ~/.config/reporter/data.db)
- Configurable Ollama host and model

## Quick start
  go 1.21+
  ollama serve            # local Ollama server
  ollama pull llama3.1    # or any model you like

  go build -o reporter ./cmd/reporter
  ./reporter add --text "Wrapped up API, fixed flaky tests" --project platform --tag api --tag tests
  ./reporter week --model llama3.1 --format markdown --per-project

## ENV / Flags
  OLLAMA_HOST   default http://127.0.0.1:11434
  --model       default "llama3.1"

## Project layout
  cmd/reporter/main.go
  internal/cmd/root.go
  internal/cmd/add.go
  internal/cmd/week.go
  internal/db/db.go
  internal/model/summary.go
  internal/ollama/client.go
  internal/summarize/prompt.go
  internal/util/date.go
