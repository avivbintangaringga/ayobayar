include .env

run: | jet-generate
	@air

jet-generate: | migrate-up
	@jet -dsn="${DATABASE_URL}" -schema=public -path=./db/jet

migrate-up:
	@goose up

migrate-down:
	@goose down

migrate-reset:
	@goose reset
