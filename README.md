# Go Mate

A Go API for dating app.

## Prerequisites

This project uses the following tech stacks:

- Go 1.23
- PostgreSQL 16

## Features

- User registration and login
- User matcher/recommendation, with pass/like functionality
- Subscriptions for extra features and boosts

## Project Structure

- `/ci` - Files related to deployment and continuous integration
- `/cmd` - Entry point `main.go` of the API.
- `/db` - SQL files for database, including migrations and seeds.
- `/internal` - Internal code for the API:
  - `/appconstant` - Constant values for business logics.
  - `/apperror` - Structured error structs for expected app errors.
  - `/config` - Configurations and settings for the app. ENV vars are loaded here.
  - `/delivery/http` - Code for handling requests and processing responses.
  - `/entity` - Base struct for business entities, directly mapped from database.
  - `/mapper` - Functions for mapping entities to response models.
  - `/model` - Base model for responses.
  - `/provider` - Functions for instantiating handler, service, repository, and utils structs
  - `/repository` - Structs for interacting with database, current implementation uses Gorm.
  - `/service` - Main app code for business logic.
  - `/util` - Utility functions and helpers.
- `/tests` - Go test files.

## Installation

For local development:

1. Copy and set env variables

    ```sh
    cp .env.example .env
    ```

2. Run migrations and seeds (if needed) using SQL files at `/db`

3. Run the app

    ```sh
    go run cmd/app/main.go
    ```

Or simply use Docker:
```sh
docker build -f ci/Dockerfile -t go-mate . && docker run -d --name go-mate go-mate
```

## Documentation

API documentation can be accessed here: [Postman docs](https://documenter.getpostman.com/view/32713619/2sAYHwK5KY)
