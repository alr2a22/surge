package logger

import (
	"os"
	"surge/internal/config"

	"github.com/sirupsen/logrus"
)

func Setup() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)
	cfg := config.GetConfig()
	if cfg.LogLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if cfg.LogLevel == "error" {
		logrus.SetLevel(logrus.ErrorLevel)
	} else {
		logrus.Infoln("Invalid log level. Log level set to info.")
	}
}
