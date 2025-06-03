# Makefile

build:
	@echo "→ Construyendo imagen…"
	@docker compose build

run-db:
	@echo "→ Iniciando contenedor de la base (db)…"
	@docker compose up -d db

run:
	@echo "→ Iniciando contenedores…"
	@docker compose up -d db
	@until docker compose exec db pg_isready -U postgres > /dev/null 2>&1; do \
	  echo "   Esperando a PostgreSQL…"; \
	  sleep 1; \
	done
	@docker compose up

migrate-up:
	@echo "→ Creando las tablas de la base de datos…"
	@goose -dir=internal/db/migrations postgres \
	  "postgres://postgres:postgres@localhost:5432/app_salud?sslmode=disable" up

migrate-down:
	@echo "→ Limpiando la base de datos (down-to 0)…"
	@goose -dir=internal/db/migrations postgres \
	  "postgres://postgres:postgres@localhost:5432/app_salud?sslmode=disable" down-to 0

test:
	@echo "→ Ejecutando tests…"
	@go test -v ./...
