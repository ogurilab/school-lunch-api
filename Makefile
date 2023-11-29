DB_URL=mysql://user:password@tcp(localhost:3306)/school_lunch?charset=utf8mb4&parseTime=True&loc=Local

TEST_DB_URL=mysql://root:root@tcp(localhost:3306)/school_lunch_test?charset=utf8mb4&parseTime=True&loc=Local

DB_URLS="$(DB_URL)" "$(TEST_DB_URL)"

APP_PATH=app

DOCKER_COMPOSE=docker compose

MIGRATION_PATH=infrastructure/db/migration

INTERFACE_SOURCES=domain/dish_domain.go domain/menu_domain.go infrastructure/db/sqlc/query.go

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
		@migrate -path $(APP_PATH)/${MIGRATION_PATH} -database "${TEST_DB_URL}" -verbose up
		@echo "\033[0;32mMigrations applied\033[0m"
	

migrateup:
	@for url in $(DB_URLS) ; do \
		migrate -path $(APP_PATH)/${MIGRATION_PATH} -database $$url -verbose up; \
	done
	@echo "\033[0;32mMigrations applied\033[0m"

migratedown:
	@for url in $(DB_URLS) ; do \
		migrate -path $(APP_PATH)/${MIGRATION_PATH} -database $$url -verbose down; \
	done
	@echo "\033[0;31mDeleted all migrations\033[0m"

new_migration:
	migrate create -ext sql -dir $(APP_PATH)/${MIGRATION_PATH} -seq $(name)

mock:	
	
	@rm -rf ${APP_PATH}/app/**/mocks 
	@echo "\033[0;31mDeleted mocks\033[0m"
	
	@cd ${APP_PATH} && for source in $(INTERFACE_SOURCES); do \
		echo "Generating mock for $$source" ; \
		dir=$$(dirname $$source); \
		filename=$$(basename $$source); \
		mockgen -source $$source -destination $$dir/mocks/$$filename -package mocks; \
	done

	@echo "\033[0;32mMocks generated\033[0m"


sqlc:
	@sqlc generate -f $(APP_PATH)/sqlc.yaml && echo "\033[0;32mSQLC generated\033[0m"


test:
	cd ${APP_PATH} && DB_SOURCE="${TEST_DB_URL}"	go test -count=1 -v -short -cover ./...


.PHONY: up down start prod prod_stop migrateup migratedown new_migration sqlc test