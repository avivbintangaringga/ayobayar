run:
	@air

migrate-up:
	@goose up

migrate-down:
	@goose down

migrate-reset:
	@goose reset
