include .env

run: | jet-generate
	@go tool air

jet-generate: | migrate-up
	@go tool jet -dsn="${DATABASE_URL}" -path=./db/jet

migrate-up:
	@go tool goose up

migrate-down:
	@go tool goose down

migrate-reset:
	@go tool goose reset
