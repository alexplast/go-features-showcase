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
    *   [Gin](https://github.com/gin-gonic/gin): A popular and performant web framework.
    *   [GORM](https://gorm.io/): A developer-friendly ORM for Go.
    *   [logrus](https://github.com/sirupsen/logrus): For structured, pluggable logging.
    *   [viper](https://github.com/spf13/viper): For application configuration management.
    *   [SQLite](https://www.sqlite.org/): A self-contained, serverless, zero-configuration, transactional SQL database engine.

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

The application is configured using a `config.yml` file. You can copy the provided `config.example.yml` to `config.yml` and modify it to your needs.

The following configuration values are available:

*   `greeting`: The greeting message to use in the application.
*   `log_level`: The logging level. Can be one of `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic`.

### Running the Application

This project uses a `Makefile` to provide convenient commands for common operations.

To start the web server, run the following command:

```sh
make run
```
You can use a tool like `curl` or Postman to interact with the API. For example:
To format the code according to Go standards, run:

```sh
make fmt
```

### Linting

This project uses `golangci-lint` for linting. To run the linter, use the following command:

```sh
make lint
```

### Testing

To run the tests, use the following command:

```sh
make test
```

## Docker

You can also build and run the application using Docker.

### Building the Docker image

To build the Docker image, run the following command:

```sh
docker build -t go-features-showcase .
```

### Running the Docker container

To run the Docker container, run the following command:

```sh
docker run -p 8080:8080 go-features-showcase
```
