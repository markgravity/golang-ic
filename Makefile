include .env
ifdef ENV
include .env.$(ENV)
endif

.PHONY: env-setup env-teardown db/migrate db/rollback migration/create migration/status dev install-dependencies test wait-for-postgres

env-setup:
	docker-compose -f docker-compose.dev.yml up -d

env-teardown:
	docker-compose -f docker-compose.dev.yml down

db/migrate:
	goose -dir database/migrations -table "migration_versions" postgres "$(DATABASE_URL)" up

db/rollback:
	goose -dir database/migrations -table "migration_versions" postgres "$(DATABASE_URL)" down

migration/create:
ifndef MIGRATION_NAME
	$(error MIGRATION_NAME is required)
endif
	goose -dir database/migrations create $(MIGRATION_NAME) sql

migration/status:
	goose -dir database/migrations -table "migration_versions" postgres "$(DATABASE_URL)" status

dev:
	make env-setup
	make db/migrate
	forego start -f Procfile.dev

install-dependencies:
	go install github.com/cosmtrek/air@v1.44.0
	go install github.com/pressly/goose/v3/cmd/goose@v3.9.0
	go install github.com/ddollar/forego@latest
	go mod tidy
	

test:
	docker-compose -f docker-compose.test.yml up -d
	go test -v -p 1 -count=1 ./...
	docker-compose -f docker-compose.test.yml down

wait-for-postgres:
	$(shell DATABASE_URL=$(DATABASE_URL) ./bin/wait-for-postgres.sh)
