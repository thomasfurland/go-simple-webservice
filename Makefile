.PHONY: test up down logs migrate-up migrate-down migrate-status

# Load env vars from test.env if it exists
ENVFILE := .env.test
ENV := $(shell [ -f $(ENVFILE) ] && cat $(ENVFILE) | xargs)

# Run all tests
test:
	$(ENV) go -C app test ./... -v

# Bring up Postgres (detached)
up:
	docker compose up -d db

# Tear down everything (remove volumes too)
down:
	docker compose down -v

# Tail logs for db
logs:
	docker compose logs -f db

migrate-up:
	$(ENV) go -C app run ./cmd/migrate

migrate-down:
	$(ENV) goose -dir ./app/migrations postgres down

migrate-status:
	$(ENV) goose -dir ./app/migrations postgres status

