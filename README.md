# Talks Indexer

A service for indexing conference talks from [moresleep](https://github.com/javaBin/moresleep) into Elasticsearch.

## Overview

Talks Indexer fetches talk/session data from a moresleep instance and bulk-indexes it into Elasticsearch. It maintains two separate indexes:

- **javazone_private**: Contains all talks with complete data, used for internal administration
- **javazone_public**: Contains only approved talks with public-safe data, used for public-facing applications

## Features

- Full reindex of all conferences, individual conferences, or single talks
- Bulk indexing for efficient Elasticsearch operations
- Dual-index strategy separating private and public data
- Simple HTTP API for triggering reindex operations
- Web admin dashboard for manual reindexing
- Webhook support for real-time updates (planned)

## Quick Start

### Prerequisites

- Go 1.22+
- Docker and Docker Compose
- Access to a running moresleep instance

### Running with Docker Compose

```bash
# Start Elasticsearch and the indexer
docker compose up -d

# Check that services are healthy
docker compose ps

# Trigger a reindex
curl -X POST http://localhost:8080/api/reindex
```

### Running Locally

```bash
# Start only Elasticsearch
docker compose up -d elasticsearch

# Set environment variables
export MORESLEEP_URL=http://localhost:8082
export ELASTICSEARCH_URL=http://localhost:9200

# Run the application
make run
```

## Configuration

Configuration is done via environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `MODE` | Running mode (`production` or `development`) | `production` |
| `PORT` | HTTP server port | `8080` |
| `MORESLEEP_URL` | Base URL of moresleep instance | `http://localhost:8082` |
| `MORESLEEP_USER` | Username for moresleep auth (optional) | - |
| `MORESLEEP_PASSWORD` | Password for moresleep auth (optional) | - |
| `ELASTICSEARCH_URL` | Elasticsearch URL | `http://localhost:9200` |
| `PRIVATE_INDEX` | Name of private index | `javazone_private` |
| `PUBLIC_INDEX` | Name of public index | `javazone_public` |

## API

### Health Check

```bash
GET /health
```

Returns service health status.

### Reindex All Conferences

```bash
POST /api/reindex
```

Triggers a full reindex of all conferences from moresleep.

### Reindex Single Conference

```bash
POST /api/reindex/conference/{slug}
```

Reindexes a specific conference by its slug (e.g., `javazone2024`).

### Reindex Single Talk

```bash
POST /api/reindex/talk/{talkId}
```

Reindexes a specific talk by its ID.

### Webhook (Planned)

```bash
POST /api/webhook
```

Receives update notifications from moresleep for incremental indexing.

## Web Admin Dashboard

A simple web interface is available at `/admin` for triggering reindex operations manually:

- Reindex all conferences
- Reindex a single conference (dropdown selection)
- Reindex a single talk (by ID)

## Architecture

The application follows hexagonal architecture principles:

```
internal/
├── adapters/           # Infrastructure implementations
│   ├── http/           # HTTP API handlers
│   ├── web/            # Web admin dashboard (templ + htmx)
│   │   ├── handlers/   # Web request handlers
│   │   └── templates/  # templ templates
│   ├── moresleep/      # Moresleep API client
│   └── elasticsearch/  # Elasticsearch client
├── app/                # Business logic
├── config/             # Configuration
├── domain/             # Domain models
└── ports/              # Interface definitions
```

## Development

```bash
# Generate templ templates
make templ

# Build (includes templ generation)
make build

# Run the application (includes templ generation)
make run

# Run tests
make test

# Format code
make fmt

# Run linter
make lint
```

## License

MIT
