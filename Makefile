DB_URL=mysql://user:password@tcp(localhost:3306)/school_lunch?charset=utf8mb4&parseTime=True&loc=Local
TEST_DB_URL=mysql://root:root@tcp(localhost:3306)/school_lunch_test?charset=utf8mb4&parseTime=True&loc=Local
DB_URLS="$(DB_URL)" "$(TEST_DB_URL)"
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

migrate_ci:
		migrate -path $(APP_PATH)/${MIGRATION_PATH} -database "${TEST_DB_URL}" -verbose up
	

migrateup:
	for url in $(DB_URLS) ; do \
		migrate -path $(APP_PATH)/${MIGRATION_PATH} -database $$url -verbose up; \
	done

migratedown:
	for url in $(DB_URLS) ; do \
		migrate -path $(APP_PATH)/${MIGRATION_PATH} -database $$url -verbose down; \
	done

new_migration:
	migrate create -ext sql -dir $(APP_PATH)/${MIGRATION_PATH} -seq $(name)

sqlc:
	sqlc generate -f $(APP_PATH)/sqlc.yaml

test:
	cd ${APP_PATH} && DB_SOURCE="${TEST_DB_URL}"	go test -v -short -cover ./...


.PHONY: up down start prod prod_stop migrateup migratedown new_migration sqlc test