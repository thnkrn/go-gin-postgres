# Restful API service with Gin, JWT, GORM, PostgresSQL and wire

## Template Structure

- [Gin](github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [JWT](github.com/golang-jwt/jwt) A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.
- [GORM](https://gorm.io/index.html) with [PostgresSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)The fantastic ORM library for Golang aims to be developer friendly.
- [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
- [fresh](https://github.com/gravityblast/fresh) is a command line tool that builds and (re)starts your web application everytime you save a Go or template file.
- [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.

## Available Scripts

In the project directory, you can run:

### `make watch`

Run and debug application with live reloading

### `make build`

Generate an execute file

### `make run`

Run an execute file

### `make serve`

Generate and run application

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
