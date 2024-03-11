include .env

run:
	go run cmd/main.go

make_migrate:
	atlas migrate diff --env gorm

migrate:
	atlas schema apply \
	--url "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL_MODE}" \
	--env gorm
