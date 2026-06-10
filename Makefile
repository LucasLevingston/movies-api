.PHONY: up down build logs test clean

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

clean:
	docker compose down -v --rmi all

health:
	curl -s http://localhost:8080/health

status:
	docker compose ps
