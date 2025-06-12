# App Salud Back

Backend para aplicación de gestión médica, escrito en Go, con PostgreSQL para almacenamiento de datos. Incluye módulos para autenticación, manejo de pacientes, citas médicas, recetas y más. Utiliza Goose para migraciones y Air para recarga en caliente durante el desarrollo.

---

## Requisitos previos

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

---

## Creacion de la base de datos

> [!IMPORTANT]
>
> Tener tu servicio de postgres detenido, en todo caso cambiar el puerto en las variables de entrorno del archivo docker-compose.yml 

1. Correr la base de datos:

```bash
make run-db
```

Equivale a:
```bash
docker compose up -d db
```

2. Levantar la base de datos:
```bash
make db-up 
```

Equivale a:
```bash
./migrate up "postgres://postgres:postgres@localhost:5432/app_salud?sslmode=disable"
```

Para Windos:
```bash 
.\migrate.exe up "postgres://postgres:postgres@localhost:5432/app_salud?sslmode=disable"
```

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
   make build
   ```

---

## Endpoints y uso básico

Ver en `http://localhost:8080/swagger/index.html`

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

