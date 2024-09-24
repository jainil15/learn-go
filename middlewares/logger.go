package middlewares

import (
	"fmt"
	"learn/go/config"
	"log/slog"
	"net/http"
	"time"
)

type RGB struct {
	R int
	G int
	B int
}

var (
	reset = "\033[0m"

	black        = RGB{0, 0, 0}
	red          = RGB{200, 100, 100}
	green        = RGB{50, 205, 50}
	yellow       = RGB{200, 200, 100}
	orange       = RGB{255, 215, 0}
	blue         = RGB{100, 100, 200}
	magenta      = RGB{200, 100, 200}
	cyan         = RGB{100, 200, 200}
	lightGray    = RGB{200, 200, 200}
	darkGray     = RGB{100, 100, 100}
	lightRed     = RGB{255, 0, 0}
	lightGreen   = RGB{0, 255, 0}
	lightYellow  = RGB{255, 255, 0}
	lightBlue    = RGB{0, 0, 255}
	lightMagenta = RGB{255, 0, 255}
	lightCyan    = RGB{0, 255, 255}
	white        = RGB{255, 255, 255}
)

type WrappedLogger struct{}

func (l *WrappedLogger) Debug(v ...interface{}) {
	env := config.Envs.Environment
	if env == "development" {
		slog.Info(colorize(fmt.Sprint(v...), lightCyan))
	}
}

type wrappedResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *wrappedResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func colorize(v string, rgb RGB) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s%s", rgb.R, rgb.G, rgb.B, v, reset)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedResponseWriter{w, http.StatusOK}
		next.ServeHTTP(wrapped, r)
		logString := ""
		env := config.Envs.Environment
		if env == "development" {
			switch {
			case wrapped.StatusCode >= 500:
				logString = colorize(fmt.Sprintf("| %v | %v | %v | %v",
					wrapped.StatusCode,
					r.Method,
					r.URL.Path,
					time.Since(start),
				), orange)
			case wrapped.StatusCode >= 400:
				logString = colorize(fmt.Sprintf("| %v | %v | %v | %v",
					wrapped.StatusCode,
					r.Method,
					r.URL.Path,
					time.Since(start),
				), orange)
			case wrapped.StatusCode >= 300:
				logString = colorize(fmt.Sprintf("| %v | %v | %v | %v",
					wrapped.StatusCode,
					r.Method,
					r.URL.Path,
					time.Since(start),
				), cyan)
			default:
				logString = colorize(fmt.Sprintf("| %v | %v | %v | %v",
					wrapped.StatusCode,
					r.Method,
					r.URL.Path,
					time.Since(start),
				), green)
			}
		} else {
			logString = fmt.Sprintf("\033[| %v | %v | %v | %v",
				wrapped.StatusCode,
				r.Method,
				r.URL.Path,
				time.Since(start),
			)
		}
		slog.Info(logString)
	})
}
