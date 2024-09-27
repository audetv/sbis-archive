init: init-ci

init-ci: docker-pull docker-build docker-up

up: docker-up
down: docker-down
restart: down up

test:
	go test -v ./...

lint:
	golangci-lint run -v

docker-up:
	docker compose up -d
docker-down:
	docker compose down --remove-orphans

docker-down-clear:
	docker compose down -v --remove-orphans
	
docker-pull:
	docker compose pull

docker-build:
	docker compose build --pull

dev-docker-build:
	REGISTRY=localhost IMAGE_TAG=main-1 make docker-build

docker-build: docker-build-service docker-build-parser
