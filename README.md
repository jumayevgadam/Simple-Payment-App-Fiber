### SIMPLE GOLANG APPLICATION GOLANG ###

✨ Features

    ⚡ Fiber v2 — Ultra-fast HTTP router inspired by Express.js

    🛡️ JWT Authentication — Secure login sessions

    📦 Modular Project Layout — Based on clean architecture

    🗃️ PostgreSQL with pgx — High-performance database access

    🧪 Validation — User input validation via go-playground/validator

    🧵 Environment Configs — Handled via envconfig and .env files

    📑 Logging — Structured logging with zap

    🌐 UUID — For consistent entity IDs

    🧰 Utility Functions — Helpful extensions with samber/lo

🧱 Tech Stack
| Layer      | Library                                                     |
| ---------- | ----------------------------------------------------------- |
| Web Server | [Fiber v2](https://github.com/gofiber/fiber)                |
| Database   | [pgx](https://github.com/jackc/pgx)                         |
| Auth       | [golang-jwt](https://github.com/golang-jwt/jwt)             |
| Validation | [validator.v10](https://github.com/go-playground/validator) |
| Logging    | [zap](https://github.com/uber-go/zap)                       |
| Config     | [envconfig](https://github.com/kelseyhightower/envconfig)   |
| UUIDs      | [google/uuid](https://github.com/google/uuid)               |

<h2> SETUP </h2>

```
    git clone https://github.com/jumayevgadam/Simple-Payment-App-Fiber.git
    cd Simple-Payment-App-Fiber
```

### CREATE ".env" file and put values ###
```
    DB_HOST = localhost
DB_PORT = 5432
DB_NAME = db_name
DB_USER = postgres
DB_PASSWORD = 12345
DB_SSLMODE = disable

## -SERVER-OPS
HTTP_PORT = 4000
SERVER_MODE = Development
READ_TIMEOUT = 10s
WRITE_TIMEOUT = 10s

## -ACCESS-OPS
JWT_SECRET_KEY = secretkey

## -LOGGER-OPS
LOG_DEVELOPMENT = true
LOG_DISABLE_CALLER = false
LOG_DISABLE_STACK_TRACE = false
LOG_ENCODING = console
LOG_LEVEL = info
```

### RUN THE APP ###

```
    go run cmd/main.go
```



🧹 Code Quality

```
    golangci-lint run ./...
```