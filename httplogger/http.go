package httplogger

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"gopkg.hlmpn.dev/pkg/go-logger"
	"gopkg.hlmpn.dev/pkg/go-logger/internal/colors"
	"gopkg.hlmpn.dev/pkg/go-logger/internal/formats"
)

var (
	ErrNilConfig = errors.New("config is nil")
)

type HTTPLogger struct {
	logger *logger.Logger
	Config *Config
}

type Config struct {
	Query     bool
	Method    bool
	JsonBody  bool
	Path      bool
	Timestamp bool
}

func NewConfig(query, method, jsonBody, path bool) *Config {
	return &Config{
		Query:    query,
		Method:   method,
		JsonBody: jsonBody,
		Path:     path,
	}
}

func New(config *Config) (*HTTPLogger, error) {

	var httpLogger *logger.Logger

	switch {
	case config == nil:
		return nil, ErrNilConfig
	case !config.Timestamp:
		httpLogger = logger.NewStdLnLogger()
	case config.Timestamp:
		httpLogger = logger.NewLogger()
	}

	return &HTTPLogger{
		logger: httpLogger,
		Config: config,
	}, nil
}

func (h *HTTPLogger) LogRequest(req *http.Request) {
	if h.Config == nil {
		return
	}

	// Extract query parameters if present
	queryParams := req.URL.RawQuery
	if queryParams == "" {
		queryParams = "-"
	}

	// Trim both path and query parameters to reasonable lengths
	trimmedPath := trimString(req.URL.Path, 100)
	trimmedQuery := trimString(queryParams, 50) //nolint:mnd //

	// Build the message with the same structure as LogRequest
	message := ""
	if h.Config.Method {
		message += colors.DarkGray + formats.Bold + req.Method + formats.Reset
	}
	message += " " + colors.DarkGray + "|" + formats.Reset + " "
	if h.Config.Path {
		message += colors.BrightBlue + trimmedPath + formats.Reset
	}
	message += " " + colors.DarkGray + "|" + formats.Reset + " "
	if h.Config.Query {
		message += colors.LightGray + trimmedQuery + formats.Reset
	}

	// Handle JSON body logging if enabled and content type is application/json
	if h.Config.JsonBody && req.Header.Get("Content-Type") == "application/json" {
		// Read the body
		body, err := io.ReadAll(req.Body)
		if err == nil {
			// Restore the body for the next handler
			req.Body = io.NopCloser(bytes.NewBuffer(body))
			// Append the JSON body to the message with a newline
			message += "\n" + colors.LightGray + string(body) + formats.Reset
		}
	}

	h.logger.Print(message)
}

// Middleware returns a handler that logs the request and then passes it to the next handler
func (h *HTTPLogger) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.LogRequest(r)
		next.ServeHTTP(w, r)
	})
}

// SetOutput sets the output writer for the logger
// Used for testing or non-stdout output
func (h *HTTPLogger) SetOutput(w io.Writer) {
	h.logger.SetOutput(w)
}
