# STJ Official (E‑commerce) Monorepo

Full‑stack e‑commerce app with a Go (Gin) backend and a React + Vite + TypeScript frontend. Backend uses PostgreSQL with sqlc for typed queries and PASETO access/refresh tokens for auth.

- Backend: `backend/` (Go 1.22, Gin, sqlc, PostgreSQL)
- Frontend: `frontend/` (React 18, Vite, Tailwind, Flowbite)
- API base path: `http://localhost:8080/api/v1`

## Contents
- Overview
- Architecture
- Local Setup
- Environment Variables
- Database & Migrations
- Run & Develop
- API Reference (auth, users, products, orders, health)
- Auth Model
- Error Format
- Testing
- Notes & Conventions

## Overview
A simple store exposing product listings, user registration/login, profile management, and order placement. The backend is versioned under `/api/v1` and uses bearer access tokens. Refresh tokens are session‑bound and persisted in the database.

## Architecture
- API server: Gin with CORS enabled (allow all origins; methods GET/POST/PUT/DELETE/OPTIONS; headers Origin, Content-Length, Content-Type, Authorization).
- Persistence: PostgreSQL. Migrations via `migrate`. Queries and models via `sqlc`.
- Auth: PASETO v2 symmetric tokens for access/refresh. Access is short‑lived; refresh is long‑lived and stored in `sessions` table.
- Versioning: `/api/v1` under `/api` prefix. Health probe at `/api/v1/healthcheck`.

## Local Setup
Prerequisites:
- Go 1.22+
- Node 18+ and Yarn or npm
- Docker (for local PostgreSQL)
- `migrate` CLI and `sqlc` (only if you need to regen SQL code)

1) Start Postgres
- Create `backend/app.env` (see sample below).
- From `backend/`: `docker compose up -d`

2) Run migrations
- From `backend/`: `make migrate_up`

3) Start backend
- From `backend/`: `go run ./...` (or `go test` to verify)

4) Start frontend
- From `frontend/`: `yarn && yarn dev` (or `npm i && npm run dev`)

## Environment Variables
Create `backend/app.env` (loaded by Viper). Example values:

```
# Database
DB_DRIVER=postgresql
DB_HOST=localhost
DB_PORT=5432
DB_NAME=stj
DB_USER=postgres
DB_PASSWORD=postgres
DB_SSL_MODE=disable

# Server
PORT=8080

# Auth (PASETO)
# Must be exactly 32 bytes (chacha20poly1305 key size)
TOKEN_SYMMETRIC_KEY=0123456789abcdef0123456789abcdef
ACCESS_TOKEN_DURATION=15m
REFRESH_TOKEN_DURATION=168h

# Admin
ADMIN_EMAIL=admin@example.com
```

Notes:
- `PORT` is propagated to Gin so `router.Run()` binds to this port.
- `TOKEN_SYMMETRIC_KEY` length must be 32 bytes. Invalid length will fail startup.

## Database & Migrations
- Migrations: `backend/db/migrations/`
- SQL queries: `backend/db/queries/`
- sqlc outputs: `backend/db/sqlc/`

Makefile commands (run from `backend/`):
- `make migrate_up` — apply all up migrations
- `make migrate_down` — roll back
- `make migrate_create DB_MIGRATION_NAME=<name>` — scaffold a migration
- `make sqlc` — regenerate sqlc code
- `make test` — run Go tests

Docker compose (Postgres) reads `backend/app.env` for DB vars.

## Run & Develop
Backend
- `cd backend`
- `go run ./...` (uses `PORT` from env, default 8080)
- Tests: `make test`

Frontend
- `cd frontend`
- `yarn && yarn dev` (Vite dev server)

CORS
- Allow all origins by default; adjust in `backend/api/server.go` if needed.

## API Reference
Base URL: `http://localhost:8080/api/v1`

Headers
- Authorization: `Bearer <access_token>` for protected endpoints
- Content-Type: `application/json`

### Health
GET `/healthcheck`
- Response 200: `{ "status": "ok" }`

### Auth
POST `/auth/register`
- Body
  - `email` (string, required)
  - `password` (string, required; base64‑encoded)
  - `language` (enum: `chn` | `jp`, required)
  - `gender` (enum: `male` | `female` | `not-specified` | `not-disclosed`)
  - Optional profile fields: `first_name`, `last_name`, `address`, `line_id`, `phone`, `latitude`, `longitude`, `birth_year`
- Response 200: user profile (see “User object” below)

POST `/auth/login`
- Body
  - `email` (string)
  - `password` (string; base64‑encoded)
- Response 200
```
{
  "access_token": "...",
  "refresh_token": "...",
  "session_id": "uuid",
  "access_token_expires": 1710000000,
  "refresh_token_expires": 1710000000,
  "user": { /* user object */ }
}
```

POST `/auth/refresh`
- Body
  - `refresh_token` (string, required)
  - `session_id` (uuid, required)
- Response 200: `{ "access_token": "..." }`

### Users (protected)
GET `/users`
- Auth: Bearer access token
- Response 200: user profile (see below)

PUT `/users`
- Auth: Bearer access token
- Body: same shape as user profile with required `email` and `language`
- Response 200: updated user profile

User object (response)
```
{
  "gender": "not-disclosed",
  "email": "user@example.com",
  "phone": "",
  "first_name": "",
  "last_name": "",
  "language": "chn",
  "address": "",
  "line_id": "",
  "latitude": 0,
  "longitude": 0,
  "birth_year": 0
}
```

### Products
GET `/products`
- Query
  - `language` (enum: `chn` | `jp`, required)
  - `limit` (int, required)
  - `offset` (int, optional; default 0)
- Response 200: array of localized product objects (joined info)

GET `/products/:id`
- Query
  - `language` (enum: `chn` | `jp`, required)
- Response 200: product with localized info and description

POST `/products` (protected)
- Auth: Bearer access token; only allowed when the authenticated user’s `email` equals `ADMIN_EMAIL` from env
- Body
```
{
  "name": "Sake X",
  "language": "chn",
  "category": "sake",
  "status": "in-stock",
  "description": {
    "introduction": "...",
    "prize": "...",
    "item_info": "...",
    "recommendation": "..."
  },
  "image_urls": ["https://..."],
  "price": 1000,
  "quantity": 5,
  "is_hot": false
}
```
- Response 200: created product summary with localized description

DELETE `/products/:id` (protected)
- Auth: Bearer access token
- Response 200: `{ "message": "product deleted" }`
- Note: Creation enforces admin check; delete currently only requires auth.

### Orders (protected)
POST `/orders`
- Auth: Bearer access token
- Body
```
{
  "shipping_address": "...",
  "phone": "...",
  "email": "buyer@example.com",
  "items": [
    { "name": "Sake X", "quantity": 2, "product_id": "uuid" }
  ]
}
```
- Response 200
```
{
  "order_id": "uuid",
  "status": "pending",
  "created_at": "...",
  "items": [ { "name": "Sake X", "quantity": 2 } ]
}
```

GET `/orders`
- Auth: Bearer access token
- Body (JSON) or Query: `limit` (int, required), `offset` (int, required)
- Response 200: array of orders with localized item names

GET `/orders/:id`
- Auth: Bearer access token
- Response 200: order with status, timestamps, and localized items

## Auth Model
- Passwords in requests are base64‑encoded on the client; backend decodes and compares against bcrypt hashes.
- Access token: short‑lived PASETO v2 token. Send via `Authorization: Bearer <token>`.
- Refresh token: long‑lived PASETO v2 token tied to a `session_id` (UUID) and stored in `sessions`. Use `/auth/refresh` to mint a new access token.
- Sessions record `user_agent`, `client_ip`, `expires_at`, and are deleted on refresh verification failure.

## Error Format
Errors are consistently returned as:
```
{ "error": "message" }
```
Common statuses: 400 (validation), 401 (unauthorized/invalid token), 403 (forbidden), 404 (not found), 409 (conflict), 500 (server/db error).

## Testing
- Backend: `cd backend && make test`
- Some tests interact with the DB; ensure Postgres and migrations are applied.

## Notes & Conventions
- API is versioned under `/api/v1`.
- CORS is permissive for local/dev. Tighten in production.
- Admin gating: product creation checks `ADMIN_EMAIL`; deletion currently only requires auth.
- Languages supported: `chn`, `jp`.
- Product and order status enums align with DB types (see `backend/db/migrations`).

## Directory Structure
- `backend/api/` — routes, handlers, middleware
- `backend/db/` — migrations, queries, sqlc generated code
- `backend/helpers/` — request helpers and error mapping
- `backend/token/` — PASETO/JWT token maker and payload
- `backend/utils/` — config, password utils
- `frontend/` — React app (pages, components, API clients)

## Development Notes
- sqlc workflow
  - Edit SQL under `backend/db/queries/`, then run `make sqlc` to regenerate types and methods in `backend/db/sqlc/`.
  - If you change enums or table shapes, add a migration first, run `make migrate_up`, then regenerate sqlc.
- Migrations
  - Create: `make migrate_create DB_MIGRATION_NAME=<snake_case_name>`
  - Apply: `make migrate_up` (ensure Postgres is running)
  - Rollback: `make migrate_down`
- Auth during development
  - Passwords sent by the frontend are base64‑encoded; the server decodes and validates against bcrypt hashes. Use HTTPS in real environments.
  - Product creation is admin‑gated by `ADMIN_EMAIL`. To create products in dev, register/login with that email.
  - Tokens: Access token in `Authorization: Bearer <token>`. Refresh requires `refresh_token` and `session_id`.
- Logging & CORS
  - Gin default logger is enabled; custom loggers are available via `server.ErrorLogger/InfoLogger/WarnLogger`.
  - CORS is wide‑open for dev in `backend/api/server.go`. Tighten for staging/prod.
- Testing
  - Backend: `make test` runs unit and integration tests. Some tests require DB and a migrated schema.
  - Prefer writing package‑level tests alongside handlers and `db/sqlc` usages.
- Frontend API base URL
  - API calls currently use `http://localhost:8080`. For configurable environments, consider refactoring to read `import.meta.env.VITE_API_BASE_URL` and define it in `frontend/.env.*` files.
- Code quality
  - Backend: `go fmt ./...`, `go vet ./...`. Optionally use `golangci-lint`.
  - Frontend: `yarn lint` (ESLint). Prettier config present (`.prettierrc.mjs`).
- Debugging tips
  - Set `GIN_MODE=debug` for verbose logs; `release` for quieter output.
  - Verify DB connectivity using `psql` against the DSN derived from env in `backend/api/server.go`.
- API evolution
  - Keep breaking changes behind a new prefix (e.g., `/api/v2`).
  - Consider adding OpenAPI/Swagger for discoverability (e.g., with `swag` for Go or a static `openapi.yaml`).
