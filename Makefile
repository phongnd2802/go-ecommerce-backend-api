GOOSE_DBSTRING ?= "root:12345@tcp(127.0.0.1:3301)/eCommerce"
GOOSE_MIGRATION_DIR ?= sql/schemas
GOOSE_DRIVER ?= mysql

APP_NAME := server

dev:
	go run ./cmd/$(APP_NAME)

docker-up:
	docker compose up -d

docker-down:
	docker compose down


create-migration:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

sqlc:
	sqlc generate

swag:
	swag init -g cmd/server/main.go -o cmd/swag/docs

.PHONY: dev docker-up docker-down upse downse sqlc create-migration swag


