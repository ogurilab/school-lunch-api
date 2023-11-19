DB_URL=mysql://user:password@localhost:3306/school_lunch
APP_PATH=app
DOCKER_COMPOSE=docker compose

dev:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yaml up -d

dev_stop:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yaml down

start:
	cd $(APP_PATH) && go run cmd/main.go

prod:
	$(DOCKER_COMPOSE) -f docker-compose.yaml up -d

prod_stop:
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
	cd ${APP_PATH} &&	go test -v -short -cover ./...


.PHONY: dev dev_stop start prod prod_stop migrateup migratedown new_migration sqlc test