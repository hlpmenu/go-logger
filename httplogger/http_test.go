package httplogger_test

import (
	"bytes"
	"encoding/json"
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	"gopkg.hlmpn.dev/pkg/go-logger/httplogger"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name     string
		query    bool
		method   bool
		jsonBody bool
		path     bool
		want     *httplogger.Config
	}{
		{
			name:     "all enabled",
			query:    true,
			method:   true,
			jsonBody: true,
			path:     true,
			want: &httplogger.Config{
				Query:    true,
				Method:   true,
				JsonBody: true,
				Path:     true,
			},
		},
		{
			name:     "all disabled",
			query:    false,
			method:   false,
			jsonBody: false,
			path:     false,
			want: &httplogger.Config{
				Query:    false,
				Method:   false,
				JsonBody: false,
				Path:     false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := httplogger.NewConfig(tt.query, tt.method, tt.jsonBody, tt.path)
			if got.Query != tt.want.Query || got.Method != tt.want.Method || got.JsonBody != tt.want.JsonBody || got.Path != tt.want.Path {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		config  *httplogger.Config
		wantErr bool
	}{
		{
			name:    "nil config",
			config:  nil,
			wantErr: true,
		},
		{
			name: "valid config without timestamp",
			config: &httplogger.Config{
				Query:    true,
				Method:   true,
				JsonBody: true,
				Path:     true,
			},
			wantErr: false,
		},
		{
			name: "valid config with timestamp",
			config: &httplogger.Config{
				Query:     true,
				Method:    true,
				JsonBody:  true,
				Path:      true,
				Timestamp: true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := httplogger.New(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("New() returned nil logger without error")
			}
		})
	}
}

func TestHTTPLogger_LogRequest(t *testing.T) {
	tests := []struct {
		name    string
		config  *httplogger.Config
		request *stdhttp.Request
	}{
		{
			name: "log method only",
			config: &httplogger.Config{
				Method: true,
			},
			request: httptest.NewRequest("GET", "/test", nil),
		},
		{
			name: "log path only",
			config: &httplogger.Config{
				Path: true,
			},
			request: httptest.NewRequest("", "/test/path", nil),
		},
		{
			name: "log query only",
			config: &httplogger.Config{
				Query: true,
			},
			request: httptest.NewRequest("", "/test?param=value", nil),
		},
		{
			name: "log json body",
			config: &httplogger.Config{
				JsonBody: true,
			},
			request: func() *stdhttp.Request {
				body := map[string]string{"test": "value"}
				jsonBody, _ := json.Marshal(body)
				req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(jsonBody))
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
		},
		{
			name: "log all fields",
			config: &httplogger.Config{
				Method:   true,
				Path:     true,
				Query:    true,
				JsonBody: true,
			},
			request: func() *stdhttp.Request {
				body := map[string]string{"test": "value"}
				jsonBody, _ := json.Marshal(body)
				req := httptest.NewRequest("POST", "/test?param=value", bytes.NewBuffer(jsonBody))
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := httplogger.New(tt.config)
			if err != nil {
				t.Fatalf("Failed to create logger: %v", err)
			}

			// Capture output
			var output bytes.Buffer
			logger.SetOutput(&output)

			// Log the request
			logger.LogRequest(tt.request)

			// Just verify we got some output
			if output.Len() == 0 {
				t.Error("LogRequest() produced no output")
			}
		})
	}
}

func TestHTTPLogger_Middleware(t *testing.T) {
	tests := []struct {
		name    string
		config  *httplogger.Config
		request *stdhttp.Request
		handler stdhttp.HandlerFunc
	}{
		{
			name: "middleware with logging",
			config: &httplogger.Config{
				Method: true,
				Path:   true,
			},
			request: httptest.NewRequest("GET", "/test", nil),
			handler: func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
				w.WriteHeader(stdhttp.StatusOK)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := httplogger.New(tt.config)
			if err != nil {
				t.Fatalf("Failed to create logger: %v", err)
			}

			// Capture output
			var output bytes.Buffer
			logger.SetOutput(&output)

			// Create test response writer
			rr := httptest.NewRecorder()

			// Create middleware handler
			handler := logger.Middleware(tt.handler)

			// Serve the request
			handler.ServeHTTP(rr, tt.request)

			// Verify we got some output
			if output.Len() == 0 {
				t.Error("Middleware() produced no output")
			}

			// Check if the response was handled
			if rr.Code != stdhttp.StatusOK {
				t.Errorf("Middleware() status code = %v, want %v", rr.Code, stdhttp.StatusOK)
			}
		})
	}
}

func TestLogRequestSimple(t *testing.T) {
	// Create test request
	req := httptest.NewRequest("GET", "/test", nil)

	// Log the request
	httplogger.LogRequestSimple(req)

	// Note: We can't capture the output in this test since it uses the global logger
	// The test will just verify that the function doesn't panic
}

func TestLogRequest(t *testing.T) {
	// Create test request
	req := httptest.NewRequest("GET", "/test?param=value", nil)

	// Log the request
	httplogger.LogRequest(req)

	// Note: We can't capture the output in this test since it uses the global logger
	// The test will just verify that the function doesn't panic
}
