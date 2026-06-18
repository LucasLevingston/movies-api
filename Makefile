.PHONY: up down build logs test test-movies test-gateway test-integration clean

up:
	docker compose up --build -d

down:
	docker compose down -v

build:
	docker compose build

logs:
	docker compose logs -f

test-movies:
	cd movies-service && go test ./... -v

test-gateway:
	cd api-gateway && go test ./... -v

test: test-movies test-gateway

test-integration:
	cd movies-service && go test -tags=integration ./internal/adapters/mongodb/... -v -timeout=120s

clean:
	docker compose down -v --rmi all

health:
	curl -s http://localhost:8080/health

status:
	docker compose ps
