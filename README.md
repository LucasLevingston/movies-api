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
├── k8s/                          # Kubernetes manifests
│   ├── namespace.yaml
│   ├── mongodb.yaml
│   ├── rabbitmq.yaml
│   ├── movies-service.yaml
│   └── api-gateway.yaml
├── movies-service/               # gRPC microservice (Hexagonal)
│   ├── internal/
│   │   ├── domain/               # Core entities
│   │   ├── ports/                # Interface definitions
│   │   ├── usecase/              # Business logic
│   │   └── adapters/
│   │       ├── grpc/             # Inbound: gRPC server
│   │       ├── mongodb/          # Outbound: MongoDB repository
│   │       ├── dynamodb/         # Outbound: DynamoDB repository (LocalStack)
│   │       └── rabbitmq/         # Outbound: event publisher
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
# Unit tests — mocks + pure domain (no external dependencies)
cd movies-service && go test ./... -v
cd api-gateway && go test ./... -v

# Integration tests — spins up a real MongoDB via testcontainers (requires Docker)
cd movies-service && go test -tags=integration ./internal/adapters/mongodb/... -v -timeout=120s
```

Or via Make:
```bash
make test              # unit tests only
make test-integration  # integration tests (requires Docker)
```

## Stopping

```bash
docker compose down -v
```

## Extras

### Event-Driven Architecture (RabbitMQ)

POST and DELETE operations publish async events via RabbitMQ. The movies-service publishes to exchanges `movie.created` and `movie.deleted` after each write.

RabbitMQ management UI is available at http://localhost:15672 (guest/guest) when running `docker compose up`.

The `RABBITMQ_URI` environment variable controls the connection (default: `amqp://guest:guest@rabbitmq:5672/`). If RabbitMQ is unreachable, the service logs a warning and continues without event publishing.

### Cloud Storage — DynamoDB via LocalStack

Switch the storage backend from MongoDB to DynamoDB (emulated locally via LocalStack) by setting `STORAGE_BACKEND=dynamodb` on the movies-service container.

```bash
# Run with DynamoDB instead of MongoDB
STORAGE_BACKEND=dynamodb docker compose up --build
```

Environment variables for the DynamoDB adapter:

| Variable              | Default                      | Description                   |
|-----------------------|------------------------------|-------------------------------|
| `STORAGE_BACKEND`     | `mongodb`                    | `mongodb` or `dynamodb`       |
| `DYNAMODB_ENDPOINT`   | *(empty — uses AWS endpoint)*| `http://localstack:4566` for LocalStack |
| `DYNAMODB_TABLE`      | `movies`                     | DynamoDB table name           |
| `AWS_REGION`          | `us-east-1`                  | AWS region                    |
| `AWS_ACCESS_KEY_ID`   | `test`                       | Credential (any value for LocalStack) |
| `AWS_SECRET_ACCESS_KEY` | `test`                     | Credential (any value for LocalStack) |

The table is created automatically on startup if it does not exist.

### Kubernetes

Manifests for all services are in `k8s/`. Apply with:

```bash
kubectl apply -f k8s/namespace.yaml
kubectl apply -f k8s/
```

This deploys: mongodb, rabbitmq, movies-service, and api-gateway into the `movies-api` namespace.

## Proto / gRPC

The `proto/movies.proto` file defines the contract between api-gateway and movies-service. Proto generation runs automatically inside the Docker build (multi-stage). To regenerate locally:

```bash
# Requires protoc + protoc-gen-go + protoc-gen-go-grpc
protoc \
  --go_out=movies-service --go_opt=module=movies-api/movies-service \
  --go-grpc_out=movies-service --go-grpc_opt=module=movies-api/movies-service \
  -I proto proto/movies.proto
```
