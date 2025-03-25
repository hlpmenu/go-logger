# Go Logger

[![Go Reference](https://pkg.go.dev/badge/gopkg.hlmpn.dev/pkg/go-logger.svg)](https://pkg.go.dev/gopkg.hlmpn.dev/pkg/go-logger)
[![Go Report Card](https://goreportcard.com/badge/gopkg.hlmpn.dev/pkg/go-logger)](https://goreportcard.com/report/gopkg.hlmpn.dev/pkg/go-logger)
[![Build Status](https://github.com/hlpmenu/go-logger/workflows/Build%20and%20Test/badge.svg)](https://github.com/hlpmenu/go-logger/actions)
[![codecov](https://codecov.io/gh/hlpmenu/go-logger/graph/badge.svg)](https://codecov.io/gh/hlpmenu/go-logger)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/hlpmenu/go-logger/badge)](https://scorecard.dev/viewer/?uri=github.com/hlpmenu/go-logger)

A simple logging package for Go applications. Log in disco colors and emojis with no setup or dependencies.


## Features

- Multiple logging levels (DEBUG, INFO, WARN, ERROR, FATAL(LogError/LogErrorf))
- Customizable log formatters
- Has aliases matching all stdlib log functions so a simple replace can be done in go.mod
- **EMOJIS** 🎉

## Installation

```bash
go get gopkg.hlmpn.dev/pkg/go-logger
```

## Basic logging

```go
package main

import (
    "gopkg.hlmpn.dev/pkg/go-logger"
)

func main() {
// Log a basic log message
log.Log("Hello, world!")

// Log a formatted log message
log.Logf("Hello, %s!", "world")

// Or the aliases
logger.Printf("Hello, %s!", "world")
logger.Print("Hello, world!")
}
```

## Log a warning message
Outputs with a yellow emoji in orange color!
```md
========================================
⚠️ Warning: [your formatted message here]
========================================
```
Usage:
```go
logger.Warnf("Something went wrong: %s", err)
// or non "f"-formatted warning message
logger.Warn("Something went wrong!")
```

## Log a info message
Outputs with a blue emoji in blue color!
```md
========================================
ℹ️ Info: [your formatted message here]
========================================
```
Usage:
```go
logger.LogInfof("Something went wrong: %s", err)
// or non "f"-formatted info message
logger.LogInfo("Something went wrong!")
```


## Log a success message
Outputs with a green emoji in green color!
```md
========================================
✅ Success: [your message here]
========================================
```
Usage:
```go
logger.LogSuccessf("Done with %s!", err)
// or non "f"-formatted success message
logger.LogSuccess("Everything is fine!")
```

## Log nice looking error messages
Outputs with a nice emoji in red color!
```md
========================================
❌ ERROR: [your formatted message here]
========================================
```
Usage:
```go
// Pre-formatted error message
logger.LogRedf("Error: %s", "Something went wrong")
// or non "f"-formatted error message
logger.LogRed("Error: Something went wrong")
```


## Log and exit using a pre-formatted error message
Output the same as above, but exits the program with a non-zero status code.

Usage:
```go
// Log and exit with a non-zero status code
logger.LogErrorf("Error: %s", "Something went wrong")
// or non "f"-formatted error message
logger.LogError("Something went wrong")
```

Note tha the above examples would have the same effect as the stdlib:
```go
log.Fatalf("Error: %s", "Something went wrong")
```
But is nicer looking and more readable.



## Documentation

For full documentation and examples, please refer to the [GoDoc](https://pkg.go.dev/gopkg.hlmpn.dev/pkg/go-logger).

## License

This project is licensed under the terms of the included LICENSE file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
