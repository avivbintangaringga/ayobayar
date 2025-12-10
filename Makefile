include .env

run:
	@air

jet-generate: | migrate-up
	@jet -dsn="${DATABASE_URL}" -schema=public -path=./db/jet -rel-model-path=../../model -rel-table-path=../../table

templ-watch:
	@templ generate -watch

tailwind-watch:
	@tailwindcss -i ./web/tailwind.css -o ./assets/static/css/app.css --watch

migrate-up:
	@goose up

migrate-down:
	@goose down

migrate-reset:
	@goose reset
