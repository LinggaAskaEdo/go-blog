package logger

import (
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once = sync.Once{}

type Options struct {
	Enabled    bool
	Path       string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func Init(opt Options) {
	once.Do(func() {
		zerolog.TimeFieldFormat = time.RFC3339Nano
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.DebugLevel)
		}

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		if opt.Enabled {
			fileLogger := &lumberjack.Logger{
				Filename:   opt.Path,
				MaxSize:    opt.MaxSize,
				MaxBackups: opt.MaxBackups,
				MaxAge:     opt.MaxAge,
				Compress:   opt.Compress,
			}

			output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
		}

		zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Int("pid", os.Getpid()).
			Logger()
	})
}
