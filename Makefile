DB_URL=mysql://user:password@localhost:3306/school_lunch
DOCKER_COMPOSE=docker compose

dev:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yaml up -d

stop-dev:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yaml down

start:
	go run cmd/main.go

prod:
	$(DOCKER_COMPOSE) -f docker-compose.yaml up -d

stop-prod:
	$(DOCKER_COMPOSE) -f docker-compose.yaml down

migrateup:
	migrate -path db/migration -database $(DB_URL) -verbose up

migratedown:
	migrate -path db/migration -database $(DB_URL) -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -short -cover ./...


.PHONY: dev stop-dev start prod migrateup migratedown new_migration sqlc test