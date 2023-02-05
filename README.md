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


-- Quais metricas? Porque Histograma, porque RED?

-- Explicar o que é white box e black box
https://sre.google/sre-book/monitoring-distributed-systems/

-- Explicar o que sao os 4 golden signals
https://sre.google/sre-book/monitoring-distributed-systems/#xref_monitoring_golden-signals

-- Diferença de RED (applicational) e USE (infrastructure) e porque usamos RED
https://blog.invgate.com/sre-signals

- Começar com o Init do instrumentation
- Fazer a parte do enpoint pra exportar metrics

- Exemplo de codigo sem template, e explicar sobre o RED

- Passagem para template, explicar um pouco mais obre o goWrap
-- Falar do generics, e da pequena gambi / martelada / solucao tatica em portugal

- Ajuste do Setup na aplicação e testar as metrics com command.

- Mostrar a quantidade de codigo escrito vs quantidade de codigo gerado...

- Falar dos comentarios.

- Dar a deixa de como automatizar esta geracao.
