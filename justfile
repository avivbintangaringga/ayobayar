set dotenv-load := true

default:
    @just --list

[parallel]
dev: watch-go watch-tailwind watch-templ

watch-go:
    air

migrate action="up":
    goose {{action}}

generate-jet: (migrate "up")
    jet -dsn="${DATABASE_URL}" -schema=public -path=./db/jet -rel-model-path=../../model -rel-table-path=../../table

watch-templ:
    templ generate --watch

watch-tailwind:
    tailwindcss -i ./web/tailwind.css -o ./assets/static/css/app.css --watch
