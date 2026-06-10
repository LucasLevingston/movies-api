# Movies API

![Tests](https://github.com/LucasLevingston/movies-api/actions/workflows/test.yml/badge.svg)

REST API built with **Go**, **Docker**, **MongoDB** and **gRPC**. Hexagonal architecture with two microservices communicating via Protocol Buffers.

## Architecture

```
┌─────────────────────────────────────────────────────────┐
│  Client (curl / browser)                                │
└────────────────────┬────────────────────────────────────┘
                     │ HTTP/REST
┌────────────────────▼────────────────────────────────────┐
│  api-gateway  (Container 1 — port 8080)                 │
│  ┌──────────────────────────────────────┐               │
│  │  Inbound:  HTTP adapter (Gin)        │               │
│  │  Domain:   Movie DTO                 │               │
│  │  Outbound: gRPC client adapter       │               │
│  └──────────────────────────────────────┘               │
└────────────────────┬────────────────────────────────────┘
                     │ gRPC / Protobuf
┌────────────────────▼────────────────────────────────────┐
│  movies-service  (Container 2 — port 50051)             │
│  ┌──────────────────────────────────────┐               │
│  │  Inbound:  gRPC server adapter       │               │
│  │  Domain:   Movie entity              │               │
│  │  Use Case: CRUD operations           │               │
│  │  Outbound: MongoDB adapter           │               │
│  └──────────────────────────────────────┘               │
└────────────────────┬────────────────────────────────────┘
                     │ MongoDB driver
┌────────────────────▼────────────────────────────────────┐
│  MongoDB  (Container 3 — port 27017)                    │
└─────────────────────────────────────────────────────────┘
```

## Prerequisites

- Docker 24+
- Docker Compose v2

## Quick Start

```bash
docker compose up --build
```

That's it. On first run the movies-service seeds the database from `movies.json` automatically.

- REST API:      http://localhost:8080
- Swagger UI:    http://localhost:8080/swagger/index.html
- Health check:  http://localhost:8080/health

## Project Structure

```
movies-api/
├── proto/                        # Shared Protobuf definitions
│   └── movies.proto
├── movies-service/               # gRPC microservice (Hexagonal)
│   ├── internal/
│   │   ├── domain/               # Core entities
│   │   ├── ports/                # Interface definitions
│   │   ├── usecase/              # Business logic
│   │   └── adapters/
│   │       ├── grpc/             # Inbound: gRPC server
│   │       └── mongodb/          # Outbound: MongoDB repository
│   └── cmd/server/main.go
└── api-gateway/                  # REST gateway (Hexagonal)
    ├── internal/
    │   ├── domain/               # DTO types
    │   ├── ports/                # Interface definitions
    │   └── adapters/
    │       ├── http/             # Inbound: Gin HTTP handlers
    │       └── grpc/             # Outbound: gRPC client
    └── cmd/server/main.go
```

## API Routes

| Method | Path            | Description            |
|--------|-----------------|------------------------|
| GET    | /movies         | List all movies        |
| GET    | /movies/:id     | Get movie by ID        |
| POST   | /movies         | Create a new movie     |
| DELETE | /movies/:id     | Delete a movie         |
| GET    | /health         | Health check           |
| GET    | /swagger/*any   | Swagger UI             |

## curl Examples

### List all movies
```bash
curl -s http://localhost:8080/movies | jq '.[0:3]'
```

### Get a movie by ID
```bash
# Get an ID from the list first
ID=$(curl -s http://localhost:8080/movies | jq -r '.[0].id')

curl -s http://localhost:8080/movies/$ID | jq
```

### Create a movie
```bash
curl -s -X POST http://localhost:8080/movies \
  -H "Content-Type: application/json" \
  -d '{"external_id": 99999, "title": "My New Movie", "year": "2024"}' | jq
```

### Delete a movie
```bash
curl -s -X DELETE http://localhost:8080/movies/$ID | jq
```

### Health check
```bash
curl -s http://localhost:8080/health
# {"status":"ok"}
```

## Running Tests

Tests require Go 1.22+ installed locally.

```bash
# movies-service (uses mocks + pure domain tests)
cd movies-service && go test ./... -v

# api-gateway (uses mocks + pure domain tests)
cd api-gateway && go test ./... -v
```

Or via Make:
```bash
make test
```

## Stopping

```bash
docker compose down -v
```

## Proto / gRPC

The `proto/movies.proto` file defines the contract between api-gateway and movies-service. Proto generation runs automatically inside the Docker build (multi-stage). To regenerate locally:

```bash
# Requires protoc + protoc-gen-go + protoc-gen-go-grpc
protoc \
  --go_out=movies-service --go_opt=module=movies-api/movies-service \
  --go-grpc_out=movies-service --go-grpc_opt=module=movies-api/movies-service \
  -I proto proto/movies.proto
```
