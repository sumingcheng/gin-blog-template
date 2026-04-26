<h1 align="center">Gin Blog Template</h1>

<p align="center">
A production-ready blog template built with Go + React.<br>
Clean architecture, dual-token auth, one-command deploy.
</p>

<p align="center">
  <a href="https://raw.githubusercontent.com/sumingcheng/gin-blog/main/LICENSE"><img src="https://img.shields.io/github/license/sumingcheng/gin-blog?color=353535" alt="license"></a>
  <a href="https://hub.docker.com/repository/docker/smcroot/gin-blog"><img src="https://img.shields.io/docker/pulls/smcroot/gin-blog?color=353535" alt="docker pull"></a>
  <a href="https://goreportcard.com/report/github.com/sumingcheng/gin-blog"><img src="https://goreportcard.com/badge/github.com/sumingcheng/gin-blog" alt="Go Report Card"></a>
</p>

<p align="center">
  <a href="./README_CN.md">中文文档</a>
</p>

<img width="3016" height="1644" alt="image" src="https://github.com/user-attachments/assets/435c1231-d30e-4919-95a1-d3ba029ae56b" />

---

## Tech Stack

| Layer | Technology |
|---|---|
| Backend | Go 1.25, Gin, GORM, PostgreSQL 16, Redis |
| Frontend | React 18, Vite, Chakra UI |
| Auth | JWT (golang-jwt/v5) + Refresh Token, bcrypt |
| Observability | Prometheus, Grafana, Logrus |
| Deploy | Docker, Docker Compose, multi-stage build |

## Project Structure

```
.
├── config/          # YAML configs (postgres, redis, jwt secret)
├── database/        # GORM models, DB/Redis connection, init.sql
├── handler/         # HTTP handlers (thin layer, delegates to service)
├── middleware/       # Auth, CORS, rate limit, request logger, metrics
├── model/           # Response structs, error codes, pagination
├── router/          # Route definitions
├── service/         # Business logic
├── util/            # JWT, bcrypt, config loader, logger, validator
├── web/             # React frontend (Vite)
├── deploy/          # Prometheus & Grafana configs
├── Dockerfile       # Multi-stage build
├── docker-compose.yaml
└── main.go
```

## API

### Public

| Method | Endpoint | Description |
|---|---|---|
| POST | `/api/register` | Register |
| POST | `/api/login` | Login (returns JWT + sets refresh token cookie) |
| POST | `/api/logout` | Logout |
| GET | `/api/token` | Refresh auth token |
| GET | `/api/blog/list` | Blog list (pagination, search) |
| GET | `/api/blog/:bid` | Blog detail |

### Authenticated

| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/user/profile` | Get profile |
| POST | `/api/user/password` | Change password |
| POST | `/api/blog/create` | Create blog |
| POST | `/api/blog/update` | Update blog |
| DELETE | `/api/blog/:bid` | Delete blog (soft delete) |

### Other

| Endpoint | Description |
|---|---|
| `/health` | Health check (DB + Redis) |
| `/metrics` | Prometheus metrics |
| `/swagger/*` | Swagger UI |

## Quick Start

### Prerequisites

- Docker & Docker Compose

### Run

```bash
git clone https://github.com/sumingcheng/gin-blog-template.git
cd gin-blog-template

# Build
make build

# Start all services
docker compose up -d
```

The app will be available at `http://localhost:5678`.

Default account: `admin` / `123456`

### Local Development

```bash
# Start PostgreSQL + Redis only
docker compose -f docker-compose-dev.yaml up -d postgres redis

# Update config/postgres.yaml and config/redis.yaml host to localhost

# Start backend
go run main.go

# Start frontend (in another terminal)
cd web && npm install && VITE_APP_ENV=development npm run dev
```

Frontend dev server runs at `http://localhost:5173`, proxying API requests to `:5678`.

## Auth Flow

```
Client                    Server                    Redis
  |                         |                         |
  |-- POST /login --------->|                         |
  |                         |-- verify password       |
  |                         |-- generate JWT          |
  |                         |-- generate refresh token|
  |                         |-- SET refresh:auth ---->|
  |<-- JWT + cookie --------|                         |
  |                         |                         |
  |-- GET /api/* ---------->|                         |
  |   (auth_token header)   |-- verify JWT            |
  |<-- response ------------|                         |
  |                         |                         |
  |-- GET /token ---------->|                         |
  |   (refresh cookie)      |-- GET refresh token --->|
  |<-- new auth_token ------|<-- auth token ----------|
```

## Monitoring

After startup, import the Grafana dashboard from `deploy/grafana/gin-blog.json`.

- Prometheus: `http://localhost:59090`
- Grafana: `http://localhost:53000` (admin / admin123456)

## License

[MIT](./LICENSE)
