package logger

import (
	"github.com/avag-sargsyan/gambling_platform/internal/conf"
	"github.com/caarlos0/env/v9"

	"github.com/rs/zerolog"
)

func Setup() {
	config := conf.Log{}
	err := env.Parse(&config)
	if err != nil {
		panic("Can't setup logger")
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	switch config.Level {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		panic("Unknown log level")
	}
}
