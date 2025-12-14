# Project Overview

This project is a Go application designed to demonstrate a wide variety of features of the Go programming language. The code is organized into a `main` package, a `features` package which contains many feature demonstrations, and a `server` package which contains a simple REST API.

This project uses the following external libraries:

*   [logrus](https://github.com/sirupsen/logrus): For structured logging.
*   [viper](https://github.com/spf13/viper): For configuration management.
*   [Gin](https://github.com/gin-gonic/gin): A popular web framework.
*   [GORM](https://gorm.io/): A developer-friendly ORM for Go.
*   [SQLite](https://www.sqlite.org/): A C-language library that implements a small, fast, self-contained, high-reliability, full-featured, SQL database engine.

# Configuration

The application is configured using a `config.yml` file. The following configuration values are available:

*   `greeting`: The greeting message to use in the application.

# Building and Running

This project uses a `Makefile` to provide convenient commands for common operations.

## Running the application

To run the program, execute the following command in your terminal:

```sh
make run
```

This will start the web server.

Alternatively, you can use the standard `go run` command:

```sh
go run main.go
```

## Building the application

To build the application binary, use the following command:

```sh
make build
```

This will create an executable file named `go-hello-world` in the root directory.

# Development Conventions

The code in this project follows standard Go formatting. You can format the code using the following command:

```sh
make fmt
```

# API Endpoints

The server provides the following endpoints for managing people:

*   `GET /people`: Get all people.
*   `POST /people`: Create a new person.
*   `GET /people/:id`: Get a person by ID.
*   `PUT /people/:id`: Update a person by ID.
*   `DELETE /people/:id`: Delete a person by ID.