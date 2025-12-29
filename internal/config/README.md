# Config Module

The config module provides centralized configuration management for the talks-indexer application using environment variables.

## Features

- Environment variable parsing using `github.com/caarlos0/env/v11`
- Optional `.env` file loading using `github.com/joho/godotenv`
- Context-based configuration access pattern
- Type-safe configuration with defaults
- Comprehensive test coverage (86.7%)

## Configuration Fields

| Field | Environment Variable | Default | Description |
|-------|---------------------|---------|-------------|
| Port | `PORT` | `8080` | HTTP server port |
| MoresleepURL | `MORESLEEP_URL` | `http://localhost:8082` | Moresleep API base URL |
| MoresleepUser | `MORESLEEP_USER` | - | Moresleep API username |
| MoresleepPassword | `MORESLEEP_PASSWORD` | - | Moresleep API password |
| ElasticsearchURL | `ELASTICSEARCH_URL` | `http://localhost:9200` | Elasticsearch connection URL |
| PrivateIndex | `PRIVATE_INDEX` | `javazone_private` | Private Elasticsearch index name |
| PublicIndex | `PUBLIC_INDEX` | `javazone_public` | Public Elasticsearch index name |

## Usage

### Basic Loading

```go
import "github.com/javaBin/talks-indexer/internal/config"

// Load configuration (returns error if parsing fails)
cfg, err := config.Load()
if err != nil {
    log.Fatal(err)
}

// Or use MustLoad for fail-fast initialization
cfg := config.MustLoad()
```

### Context-Based Access

```go
import (
    "context"
    "github.com/javaBin/talks-indexer/internal/config"
)

func main() {
    cfg := config.MustLoad()
    ctx := config.WithConfig(context.Background(), cfg)

    // Pass context through your application
    doWork(ctx)
}

func doWork(ctx context.Context) {
    cfg := config.GetConfig(ctx)
    fmt.Printf("Server port: %d\n", cfg.Port)
}
```

### Environment Variables

Create a `.env` file in the project root (optional):

```env
PORT=8080
MORESLEEP_URL=https://api.example.com
MORESLEEP_USER=myuser
MORESLEEP_PASSWORD=mypassword
ELASTICSEARCH_URL=http://localhost:9200
PRIVATE_INDEX=javazone_private
PUBLIC_INDEX=javazone_public
```

Or export environment variables directly:

```bash
export PORT=8080
export MORESLEEP_URL=https://api.example.com
```

## Testing

Run tests with coverage:

```bash
make test
make coverage
```

The module has 86.7% test coverage, exceeding the 75% minimum requirement.

## Architecture

This module follows clean architecture principles:

- Self-contained with no dependencies on other internal packages
- Used by adapters to access their configuration
- Configuration is injected via context throughout the application
- Fails fast on initialization errors using `MustLoad()`
