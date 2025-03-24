# Go Logger

A robust, flexible logging package for Go applications, providing various logging levels, formatters, and output destinations.

## Features

- Multiple logging levels (DEBUG, INFO, WARN, ERROR, FATAL)
- Customizable log formatters
- Various output destinations (console, file, network)
- Context-based logging
- Structured logging support
- Thread-safe operations

## Installation

```bash
go get github.com/username/go-logger
```

## Quick Start

```go
package main

import (
    "github.com/username/go-logger"
)

func main() {
    // Initialize a default logger
    log := logger.New()
    
    // Log messages at different levels
    log.Debug("Debug message")
    log.Info("Information message")
    log.Warn("Warning message")
    log.Error("Error message")
    
    // With context data
    log.WithField("user_id", "12345").Info("User logged in")
}
```

## Subpackages

- **errutil**: Error handling utilities for logging
- **sysutils**: System utilities for logging operations

## Documentation

For full documentation and examples, please refer to the [GoDoc](https://godoc.org/github.com/username/go-logger).

## License

This project is licensed under the terms of the included LICENSE file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
