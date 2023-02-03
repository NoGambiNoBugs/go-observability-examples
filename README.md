# go-observability-examples

Example of observability in golang.

## Setup

```bash
# CREATE .env with following envs
SERVER_ADDRESS=":8080"
DB_USER="example-local"
DB_PASSWORD="example-local"
DB_DATABASE="example-local"
DB_ADDRESS="localhost"
DB_PORT="5432"
DB_MIGRATIONS_PATH="./infra/migrations"
DB_EXTENSIONS_FILEPATH="./infra/scripts/extensions.sql"

# UP infrastructure
docker-compose up -d

# Create extensions
go run ./cmd/cli/cli.go extensions

# Migrate tables
go run ./cmd/cli/cli.go migrate

# Run Server API
go run ./cmd/server/server.go
```

## Testing

```curl

```

## Shutdown

```bash
# DOWN infrastructure
docker-compose down
```
