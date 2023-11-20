DB_URL=mysql://user:password@tcp(localhost:3306)/school_lunch?charset=utf8mb4&parseTime=True&loc=Local
APP_PATH=app
DOCKER_COMPOSE=docker compose
MIGRATION_PATH=infrastructure/db/migration

up:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yaml up -d

down:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yaml down

start:
	cd $(APP_PATH) && go run cmd/main.go

prod:
	$(DOCKER_COMPOSE) -f docker-compose.yaml up -d

prod_stop:
	$(DOCKER_COMPOSE) -f docker-compose.yaml down

migrateup:
	migrate -path $(APP_PATH)/${MIGRATION_PATH} -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path $(APP_PATH)/${MIGRATION_PATH} -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir $(APP_PATH)/${MIGRATION_PATH} -seq $(name)

sqlc:
	sqlc generate -f $(APP_PATH)/sqlc.yaml

test:
	cd ${APP_PATH} &&	go test -v -short -cover ./...


.PHONY: up down start prod prod_stop migrateup migratedown new_migration sqlc test