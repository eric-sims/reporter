# reporter — daily summaries CLI (projects & tags)

A Go CLI to capture daily summaries and generate weekly summaries, grouped by project, using a locally running Ollama LLM.

## Features
- `reporter add` save a summary for today or a specific date (from flag, stdin, or $EDITOR)
- `reporter week` aggregate a given week (or current week) and ask Ollama to summarize
- `reporter list` lists reports for a given day (default today)
- SQLite storage via GORM (DB file: ~/.config/reporter/data.db)
- Configurable Ollama host and model

## Installation
TODO

## Connect Ollama
TODO

## Upcoming features
- [ ] Ability to connect to ChatGPT API
