# App Salud Back

Backend para aplicación de gestión médica, escrito en Go, con PostgreSQL para almacenamiento de datos. Incluye módulos para autenticación, manejo de pacientes, citas médicas, recetas y más. Utiliza Goose para migraciones y Air para recarga en caliente durante el desarrollo.

---

## Tabla de contenidos

1. [Características principales](#caracter%C3%ADsticas-principales)  
2. [Requisitos previos](#requisitos-previos)  
3. [Configuración](#configuraci%C3%B3n)  
4. [Migraciones de base de datos](#migraciones-de-base-de-datos)  
5. [Ejecución en desarrollo (con Air)](#ejecuci%C3%B3n-en-desarrollo-con-air)  
6. [Ejecución con Docker Compose](#ejecuci%C3%B3n-con-docker-compose)  
7. [Endpoints y uso básico](#endpoints-y-uso-b%C3%A1sico)  
8. [Testing](#testing)  
9. [Variables de entorno](#variables-de-entorno)  
11. [Licencia](#licencia)  

---

## Características principales

- **Autenticación** (JWT con HS256)  
- **Gestión de usuarios** (registro, login, roles)  
- **CRUD de pacientes**  
- **CRUD de citas médicas**  
- **CRUD de diagnósticos, tratamientos y resultados de laboratorio**  
- **CRUD de vacunas y antecedentes**  
- **CRUD de recetas y sus items**  
- **Middleware de validación y autorización**  
- **Migraciones con Goose**  
- **HotReloading con Air**  

---

## Requisitos previos

- Go 1.24+  
- Docker + Docker Compose (si se desea ejecutar con contenedores)  
- `make` (opcional, para comandos definidos en Makefile)  
- Cliente HTTP (por ejemplo, Postman, Insomnia o CURL)  


---

## Configuración

1. Clona el repositorio:
   ```bash
   git clone <url-del-repo> app-salud-back
   cd app-salud-back
   ```

2. Crea un archivo `.env` en la raíz (o exporta variables de entorno manualmente). Variables comunes:

   ```env
   # Para la app Go
   DATABASE_URL=postgres://postgres:postgres@db:5432/app_salud?sslmode=disable
   PORT=8080
   TOKEN_SECRET=secreto-para-jwt
   ```

   > Si ejecutas con Docker Compose, se usan directamente las variables definidas en `docker-compose.yml`.

3. Instala las dependencias Go (ya incluidas en go.mod/go.sum):
   ```bash
   go mod download
   ```

---

## Migraciones de base de datos

Se usan migraciones con Goose para elevar/escalar y revertir cambios en la base de datos.

- **Aplicar todas las migraciones pendientes**:
  ```bash
  make migrate-up
  ```
  Internamente hace:
  ```bash
  goose -dir=internal/db/migrations postgres "$DATABASE_URL" up
  ```

- **Revertir todas las migraciones (down-to 0)**:
  ```bash
  make migrate-down
  ```
  Equivale a:
  ```bash
  goose -dir=internal/db/migrations postgres "$DATABASE_URL" down-to 0
  ```

- **Verificar el estado de migraciones**:
  ```bash
  goose -dir=internal/db/migrations postgres "$DATABASE_URL" status
  ```

> Asegúrate de que la variable `DATABASE_URL` apunte al contenedor PostgreSQL (cuando corras con Docker, suele ser `postgres://postgres:postgres@db:5432/app_salud?sslmode=disable`).

---

## Ejecución en desarrollo (con Air)

Air permite recarga en caliente de la aplicación Go cada vez que guardas cambios.

1. Instala Air (si no está instalado):
   ```bash
   go install github.com/air-verse/air@latest
   ```
2. Ejecuta Air:
   ```bash
   air -c .air.toml
   ```
   Esto compilará y levantará el servidor en `localhost:8080`. Cada vez que guardes un archivo `.go`, Air reconstruirá y reiniciará automáticamente.

3. Verás logs nivel DEBUG de cambios (según `.air.toml`).

---

## Ejecución con Docker Compose

Para levantar la base de datos PostgreSQL junto con el servidor Go en modo desarrollo:

1. Levanta los contenedores (modo background):
   ```bash
   make run
   ```
   Esto hace:
   - `docker compose up -d db` → crea y arranca solo el contenedor `db`.  
   - Espera a que PostgreSQL esté listo (`pg_isready`).  
   - `docker compose up` → inicia el contenedor `api` (Go + Air) enlazado al `db`.

2. Para volver a construir la imagen si modificas el `Dockerfile`:
   ```bash
   docker compose build
   ```

3. Verifica que `db` esté corriendo en el puerto 5432 y `api` en el puerto 8080.  
   - Ejemplo: `http://localhost:8080/health` (endpoint de prueba, si lo definiste).

4. Para bajar todo:
   ```bash
   docker compose down
   ```

5. Si necesitas eliminar datos de la base (volumen):
   ```bash
   docker compose down -v
   ```

---

## Endpoints y uso básico

Ver en `http://localhost:8080/swagger/index.html`

--- 

## Testing

Para ejecutar pruebas unitarias:

```bash
make test
```

o directamente:

```bash
go test -v ./...
```

---

## Variables de entorno

- `DATABASE_URL`  
  URL de conexión a PostgreSQL. Ejemplo:  
  ```
  postgres://postgres:postgres@db:5432/app_salud?sslmode=disable
  ```
- `PORT`  
  Puerto en el que corre la API (por defecto `8080`).
- `TOKEN_SECRET`  
  Secreto HMAC para firmar/verificar JWT (si usas HS256).

> Si usas ES256, en lugar de `TOKEN_SECRET` deberás cargar las claves ECDSA de un archivo PEM (modificar código en `internal/utils/jwt.go`).

---

## Licencia

Este proyecto está bajo la **MIT License**.  
