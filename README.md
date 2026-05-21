# boteco

CLI chat assistant in Go. Powered by [Genkit](https://github.com/firebase/genkit) + OpenAI (`gpt-5-nano`). Streams replies in the terminal and can call tools.

## Tools

- `web_search` — DuckDuckGo HTML scrape for time-sensitive queries.
- `create_event` — insert event (description, date) into local SQLite.
- `fetch_events` — list stored events.

## Stack

- Go 1.26
- `github.com/firebase/genkit/go`
- `modernc.org/sqlite` (pure-Go SQLite)

## Setup

```sh
cp .env.example .env
# edit .env:
#   OPENAI_API_KEY=...
#   DB_PATH=./boteco.db
```

## Run

```sh
task dev
# or
go run main.go
```

Type at the `>` prompt. `Ctrl+C` or EOF to exit.
