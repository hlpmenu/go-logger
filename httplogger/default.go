package httplogger

import (
	"net/http"

	"gopkg.hlmpn.dev/pkg/go-logger"
	"gopkg.hlmpn.dev/pkg/go-logger/internal/colors"
	"gopkg.hlmpn.dev/pkg/go-logger/internal/formats"
)

func LogRequestSimple(req *http.Request) {
	logger.Logf("Request: %s %s", req.Method, req.URL)
}

func LogRequest(req *http.Request) {
	logRequest(req)
}

// logRequest logs the HTTP request details in a Gin-style format
func logRequest(r *http.Request) {
	// Extract query parameters if present
	queryParams := r.URL.RawQuery
	if queryParams == "" {
		queryParams = "-"
	}

	// Trim both path and query parameters to reasonable lengths
	trimmedPath := trimString(r.URL.Path, 100)
	trimmedQuery := trimString(queryParams, 50) //nolint:mnd //

	logger.Printf("%s%s%s%s %s|%s %s%s%s %s|%s %s%s%s",
		colors.DarkGray,
		formats.Bold,
		r.Method,
		formats.Reset,
		colors.DarkGray,
		formats.Reset,
		colors.BrightBlue,
		trimmedPath,
		formats.Reset,
		colors.DarkGray,
		formats.Reset,
		colors.LightGray,
		trimmedQuery,
		formats.Reset)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		LogRequest(r)
		next.ServeHTTP(w, r)
	})
}

// colorStatusCode returns the appropriate color code based on status code range
// func colorStatusCode(statusCode int) string {
// 	switch {
// 	case statusCode >= 200 && statusCode < 300:
// 		return colors.Green
// 	case statusCode >= 300 && statusCode < 400:
// 		return colors.Cyan
// 	case statusCode >= 400 && statusCode < 500:
// 		return colors.Yellow
// 	default:
// 		return colors.Red
// 	}
// }

// trimString shortens a string to the specified length, adding "..." if trimmed
func trimString(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen-3] + "..."
	}
	return s
}
