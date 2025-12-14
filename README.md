# Go Hello World - Feature Demonstration

This project is a Go application designed to demonstrate a wide variety of features of the Go programming language. It serves as a comprehensive example of how to use various libraries and language constructs to build a modern Go application.

## Features

This project demonstrates the following Go features and libraries:

*   **Language Constructs:**
    *   Structs and Methods
    *   Interfaces
    *   Goroutines and Channels
    *   Error Handling
    *   The `defer` statement
    *   Slices and Maps
    *   Pointers
    *   The `select` statement for concurrent operations
    *   `sync.Mutex` for thread-safe code

*   **Standard Library:**
    *   `encoding/json` for JSON manipulation
    *   `net/http` for making HTTP requests
    *   `context` for managing request-scoped values, cancellation signals, and deadlines

*   **External Libraries:**
    *   [logrus](https://github.com/sirupsen/logrus): For structured logging.
    *   [viper](https://github.com/spf13/viper): For configuration management.
    *   [Gin](https://github.com/gin-gonic/gin): A popular web framework.
    *   [GORM](https://gorm.io/): A developer-friendly ORM for Go.
    *   [SQLite](https://www.sqlite.org/): A C-language library that implements a small, fast, self-contained, high-reliability, full-featured, SQL database engine.
    *   [stretchr/testify](https://github.com/stretchr/testify): For testing utilities.

## Getting Started

### Prerequisites

*   Go (version 1.18 or later)
*   Make

### Installation

1.  Clone the repository:
    ```sh
    git clone https://github.com/your-username/go-features-showcase.git
    cd go-features-showcase
    ```

2.  Install the dependencies:
    ```sh
    go mod tidy
    ```

### Configuration

The application is configured using a `config.yml` file. The following configuration values are available:

*   `greeting`: The greeting message to use in the application.
*   `log_level`: The logging level. Can be one of `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic`.

## Building and Running

This project uses a `Makefile` to provide convenient commands for common operations.

### Running the application

To run the program, execute the following command in your terminal:

```sh
make run
```

This will start the web server with graceful shutdown.

Alternatively, you can use the standard `go run` command:

```sh
go run main.go
```

### Building the application

To build the application binary, use the following command:

```sh
make build
```

This will create an executable file named `go-features-showcase` in the root directory.

### Running the demo

To run a demonstration of the application's functionality, use the following command:

```sh
make demo
```

This will start the server, make a request to the `/people` endpoint, and then stop the server.

### Running with Docker

You can also build and run the application using Docker.

#### Building the Docker image

To build the Docker image, run the following command:

```sh
docker build -t go-features-showcase .
```

#### Running the Docker container

To run the Docker container, run the following command:

```sh
docker run -p 8080:8080 go-features-showcase
```

# Development Conventions

The code in this project follows standard Go formatting, linting, and testing practices.

## Formatting

You can format the code using the following command:

```sh
make fmt
```

This uses `gofumpt` for stricter formatting.

## Linting

This project uses `golangci-lint` for linting. To run the linter, use the following command:

```sh
make lint
```

## Testing

To run the tests, use the following command:

```sh
make test
```

# API Endpoints

The server provides the following endpoints for managing people:

*   `GET /people`: Get all people.
*   `POST /people`: Create a new person.
*   `GET /people/:id`: Get a person by ID.
*   `PUT /people/:id`: Update a person by ID.
*   `DELETE /people/:id`: Delete a person by ID.

