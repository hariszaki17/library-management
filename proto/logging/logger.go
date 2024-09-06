package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is the global logger instance
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	// Set the output to stdout
	Logger.SetOutput(os.Stdout)

	// Set the log level (could be logrus.DebugLevel, logrus.InfoLevel, etc.)
	Logger.SetLevel(logrus.DebugLevel)

	// Set the formatter to JSON
	Logger.SetFormatter(&logrus.JSONFormatter{})
}
