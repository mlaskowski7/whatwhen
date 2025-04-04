# Setup

## Db setup

Create `db.env` file in the root dir and fill it according to the example:
```
POSTGRES_USER: postgres
POSTGRES_PASSWORD: postgres
POSTGRES_DB: whatwhen
```

## App setup

Create `app.env` file in the root dir and fill it according to the example:
```
DB_HOST=whatwhen_db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=whatwhen
```

# Usage

Create an cli alias for this app
```bash
alias whatwhen='sh -c '\''docker compose run --rm whatwhen_app "$@" 2>&1 | grep -vE "^\[\+|^\s+✔|^\s+Container "'\'' --'
source ~/.bashrc  # or ~/.zshrc
```

Run the compose before executing any command
```bash
docker compose up -d
```

# Commands

## Whats

How to use:
```bash
whatwhen whats
```

## New

How to use:
```bash
whatwhen new \
  --title "Deploy app" \
  --desc "Finish container setup" \
  --deadline "2025-04-10T20:00" \
  --finished=false
```