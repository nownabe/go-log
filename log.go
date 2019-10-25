package log // import "go.nownabe.dev/log"

import "go.uber.org/zap"

type wfunc func(string, ...interface{})

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger

	// Debugw logs a message with some additional context.
	// The variadic key-value pairs are treated as they are in With.
	Debugw wfunc

	// Infow logs a message with some additional context.
	// The variadic key-value pairs are treated as they are in With.
	Infow wfunc

	// Warnw logs a message with some additional context.
	// The variadic key-value pairs are treated as they are in With.
	Warnw wfunc

	// Errorw logs a message with some additional context.
	// The variadic key-value pairs are treated as they are in With.
	Errorw wfunc

	// Panicw logs a message with some additional context,
	// then panics. The variadic key-value pairs are treated as they are in With.
	Panicw wfunc

	// Sync flushes any buffered log entries.
	Sync func() error
)

// Init initializes logger.
// If development is true,
// it builds development logger that writes Debug and above logs
// to stderr in a human-friendly format.
func Init(development bool) {
	var err error

	if development {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	sugar = logger.Sugar()

	Debugw = sugar.Debugw
	Infow = sugar.Infow
	Warnw = sugar.Warnw
	Errorw = sugar.Errorw
	Panicw = sugar.Panicw
	Sync = sugar.Sync
}

// Logw logs a message as specified log level.
func Logw(l Level, msg string, keysAndValues ...interface{}) {
	switch l {
	case LevelDebug:
		Debugw(msg, keysAndValues...)
	case LevelInfo:
		Infow(msg, keysAndValues...)
	case LevelWarn:
		Warnw(msg, keysAndValues...)
	case LevelError:
		Errorw(msg, keysAndValues...)
	}
}
