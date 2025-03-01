include .env
export $(shell sed -n 's/=.*//p' .env)

migrate_up:
	migrate -path db/migration -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up