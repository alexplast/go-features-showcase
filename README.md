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

### Running the Application

This project uses a `Makefile` to provide convenient commands for common operations.

To start the web server, run the following command:

```sh
make run
```

The server will start on `http://localhost:8080`.

## API Endpoints

The server provides the following endpoints for managing people:

| Method | Endpoint         | Description          |
| ------ | ---------------- | -------------------- |
| `GET`    | `/people`        | Get all people       |
| `POST`   | `/people`        | Create a new person  |
| `GET`    | `/people/:id`    | Get a person by ID   |
| `PUT`    | `/people/:id`    | Update a person by ID|
| `DELETE` | `/people/:id`    | Delete a person by ID|

You can use a tool like `curl` or Postman to interact with the API. For example:

```sh
# Get all people
curl http://localhost:8080/people

# Create a new person
curl -X POST http://localhost:8080/people -H "Content-Type: application/json" -d '{"name": "Jane", "age": 25}'
```

## Development

### Building

To build the application binary, use the following command:

```sh
make build
```

This will create an executable file named `go-features-showcase` in the root directory.

### Formatting

To format the code according to Go standards, run:

```sh
make fmt
```
