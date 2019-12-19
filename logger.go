package micro

import jaeger "github.com/uber/jaeger-client-go"

var logger jaeger.Logger

func init() {
	logger = jaeger.NullLogger
}

// SetLogger sets the logger
func SetLogger(l jaeger.Logger) {
	logger = l
}

// Logger gets the logger
func Logger() jaeger.Logger {
	return logger
}
