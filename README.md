# MemeCoin API

## Table of Contents

- [File Structure](#file-structure)
- [Environment Variables](#environment-variables)
- [Installation](#installation)
- [Usage](#usage)

## File Structure

```
/portto-assignment
├── api/
├── assets/
│   └── sql/
├── cmd/
├── config/
├── internal/
| ├── handlers/
| ├── repositories/
| ├── routes/
| └── services/
├── scripts/
├── test/
| └── mocks/
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

## Environment variables

本專案使用以下環境變數來配置 PostgreSQL 和服務環境。

### Docker Compose - PostgreSQL 設定

使用 Docker Compose 啟動 PostgreSQL 時，需要設定以下設定：

| Environment variables | 說明                        |
| --------------------- | --------------------------- |
| `POSTGRES_USER`       | PostgreSQL 初始化用戶名     |
| `POSTGRES_PASSWORD`   | PostgreSQL 初始化密碼       |
| `POSTGRES_DB`         | PostgreSQL 初始化數據庫名稱 |
| `REDIS_PASSWORD`      | Redis 初始化密碼            |

### Application

Application 本身需要以下設定：

| Environment variables | 說明                                            |
| --------------------- | ----------------------------------------------- |
| `POSTGRESQL_URL`      | Application 使用的 PostgreSQL connection string |
| `REDIS_URL`           | Application 使用的 Redis connection string      |

### 環境設定方式

- **本地開發環境**：使用 `./config` 中的 `config.env.local`
- **其他環境**：需手動設定

### Example

#### Docker Compose PostgreSQL 配置

`./.env`

```env
POSTGRES_USER="admin"
POSTGRES_PASSWORD="password123"
POSTGRES_DB="myapp"
REDIS_PASSWORD="redispass123"
```

#### 應用程式配置

`./config/config.env.local`

```env
POSTGRESQL_URL="postgresql://admin:password123@localhost:5432/myapp"
REDIS_URL="redis://redispass123@localhost:6379/0?protocol=3"

```

## Installation

```bash
# Clone the repository
git clone https://github.com/frank00125/portto-assignment.git

# Navigate to the project directory
cd portto-assignment

# Install dependencies
go get

# Set up the env for docker compose
# You need to fill in the environment variables to .env file
cat .env.example > .env

# Set up local directory to put the database files
mkdir docker-database

# Start docker compose at the background
docker compose up -d

# Setup the env for application
# You need to fill in the environment variables to ./config/config.env.local file
cat ./config/config.env.example > ./config/config.env.local

# Database seeding
go run ./seeds/seeds.go

# Start the project
go run ./cmd/main.go
```

## Usage

於 local 環境啟動 dev server

```bash
# Start the dev server
go run ./cmd/main.go
```

執行單元測試

```bash
# Run test (without cache)
go clean -testcache && go test -v ./...
```

更新 API 文件

```bash
# Generate new API documentation after development
swag init --g ./cmd/main.go --output ./api --outputTypes yaml
```
