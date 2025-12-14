# Project Overview

This project is a Go application designed to demonstrate a wide variety of features of the Go programming language. The code is organized into a `main` package and a `features` package, which contains the demonstration functions.

This project uses the following external libraries:

*   [logrus](https://github.com/sirupsen/logrus): For structured logging.
*   [viper](https://github.com/spf13/viper): For configuration management.

The demonstrated features include:

*   Structs and Methods
*   Interfaces
*   Goroutines and Channels
*   JSON Marshaling
*   Error Handling
*   The `defer` statement
*   Slices
*   Maps
*   Pointers
*   The `select` statement
*   `sync.Mutex` for thread safety
*   String manipulation
*   File I/O
*   `context` for cancellation
*   `net/http` for making HTTP requests

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
