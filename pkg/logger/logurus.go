package logger

import (
	"dagger-pac/config"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"os"
)

func LogurusSetup(cfg *config.Config) Interface {
	// select which logger to use after reading the config
	return &log.Logger{
		Out:   os.Stderr,
		Level: getLogLevel(cfg.Log.Level),
		Formatter: &nested.Formatter{
			HideKeys:    true,
			CallerFirst: true,
		},
	}
}

func getLogLevel(level string) log.Level {
	switch level {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	case "panic":
		return log.PanicLevel
	default:
		return log.PanicLevel
	}
}
