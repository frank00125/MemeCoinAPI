# Project Title

A brief description of what this project does and who it's for.

## Table of Contents

- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [Installation](#installation)
- [Usage](#usage)
- [Docker](#docker)

## Project Structure

An overview of the project's structure and layout.

```
/portto-assignment
├── config/
├── handlers/
├── mocks/
├── repositories/
├── routes/
├── seeds/
├── services/
├── static/
│   └── sql/
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Environment variables

We use the following environment to start this project

- SERVICE_ENV: The current environment
- POSTGRESQL_DATABASE_URL: The connection string used in the application
- POSTGRES_USER: The initialized username for starting PostgreSQL for docker compose usage
- POSTGRES_PASSWORD: The initialized password for starting PostgreSQL for docker compose usage
- POSTGRES_DB: The initialized database name for starting PostgreSQL for docker compose usage

## Installation

Steps to set up and configure the project.

```bash
# Clone the repository
git clone https://github.com/frank00125/portto-assignment.git

# Navigate to the project directory
cd portto-assignment

# Install dependencies
go get

# Start docker compose at the background
docker compose up -d

# Setup the env
# You need to fill in the environment variables to .env file
cat .env.example > .env

# Database seeding
go run seeds/seeds.go
```

## Usage

Instructions for running the application locally.

```bash
# Start the dev server
go run main.go
```

Instructions for running the unit test

```bash
# Run test (without cache)
go clean -testcache && go test -v ./...
```

Instructions for generating new API documentation

```bash
# Generate new API documentation after development
swag init --outputTypes yaml
```

## Docker

We use `docker compose` to start local database for the development. Database files will be in the `/docker-database` directory

```

```
