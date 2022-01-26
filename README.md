# Restful API service with Gin, JWT, GORM, PostgresSQL and wire

## Template Structure

- [Gin](github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [JWT](github.com/golang-jwt/jwt) A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.
- [GORM](https://gorm.io/index.html) with [PostgresSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)The fantastic ORM library for Golang aims to be developer friendly.
- [WIRE](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.

## Available Scripts

In the project directory, you can run:

### `go build`

Generate an execute file

### `./go-gin-postgres`

Run project

## Available Endpoint

In the project directory, you can call:

### `GET /`

For getting status page

### `POST /login`

For generating a JWT

### `GET /api/campaigns`

For getting all of campaigns

### `GET /api/campaigns/:id`

For getting campaign by ID

### `POST /api/campaigns`

For creating new campaign

### `DELETE /api/campaigns/:id`

For removing existing campaign

### `PUT /api/campaigns/:id`

For updating existing campaign
