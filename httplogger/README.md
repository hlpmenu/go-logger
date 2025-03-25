[![Go Reference](https://pkg.go.dev/badge/gopkg.hlmpn.dev/pkg/go-logger/httplogger.svg)](https://pkg.go.dev/gopkg.hlmpn.dev/pkg/go-logger/httplogger)
[![Go Report Card](https://goreportcard.com/badge/gopkg.hlmpn.dev/pkg/go-logger/httplogger)](https://goreportcard.com/report/gopkg.hlmpn.dev/pkg/go-logger/httplogger)

# HTTP Logger

A simple HTTP logger for the stdlib net/http package.


## Usage

```go
package main

import (
	"net/http"

	"gopkg.hlmpn.dev/pkg/go-logger/httplogger"
)

func main() {
	// Create a new HTTP logger
	logger := httplogger.New(httplogger.NewConfig(true, true, true, true))

	// Create a new HTTP server
	server := http.Server{
		Addr:    ":8080",
		Handler: logger.Middleware(http.HandlerFunc(handler)),
	}

	// Start the server
	if err := server.ListenAndServe(); err != nil {
		logger.LogError(err)
	}
}
```

## Middleware

The `Middleware` function is a convenient way to add the HTTP logger to an existing HTTP server.

```go
http.DefaultServeMux.Handle("/", logger.Middleware(http.HandlerFunc(handler)))
```

## Configuration(Optional)

The `Config` struct allows you to customize the HTTP logger's behavior.

```go
type Config struct {
	Query     bool // Log query parameters in the request URL
	Method    bool // Log the HTTP method
	JsonBody  bool // Log the JSON body nicely formatted. Works only with Content-Type: application/json
	Path      bool // Log the request path
	Timestamp bool // Log with a timestamp(the standard timestamp format, not a extra one for the request)
}
```

## 0-config logger
```go
handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	httplogger.LogRequest(r)
	w.WriteHeader(http.StatusOK)
})
```

## 0-config middleware
```go
http.DefaultServeMux.Handle("/", httplogger.LogRequest(http.HandlerFunc({
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, world!"))
})))
```

## Only log method, path 
```go
handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	httplogger.LogRequestSimple(r)
	w.WriteHeader(http.StatusOK)
})
```


