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

db-up:
	@echo "→ Levantando la base de datos…"
	@./migrate up "postgres://postgres:postgres@localhost:5432/app_salud?sslmode=disable"

db-down:
	@echo "→ Limpiando la base de datos (down-to 0)…"
	@./migrate down-to "postgres://postgres:postgres@localhost:5432/app_salud?sslmode=disable" 0
