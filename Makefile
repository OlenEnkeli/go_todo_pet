include .env

SHELL=bash

name = ''
postgres_uri = "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL_MODE}"

run:
	go run cmd/main.go

make_migration:
	migrate create -ext sql -dir ./migrations -seq ${name}

migration_up:
	migrate -path ./migrations -database ${postgres_uri} up

migration_down:
	migrate -path ./migrations -database ${postgres_uri} down 1
