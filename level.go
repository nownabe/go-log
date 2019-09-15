package log

// Level is a logging priority.
// Higher levels are more important.
type Level int

const (
	// LevelDebug logs are typically voluminous,
	// and are usually disabled in production.
	LevelDebug Level = 1
	// LevelInfo is the default logging priority.
	LevelInfo Level = 2
	// LevelWarn logs are more important than Info,
	// but donn't need individual human review.
	LevelWarn Level = 3
	// LevelError logs are high-priority.
	// If an application is running smoothly,
	// Error logs shouldn't appear.
	LevelError Level = 4
)

// String returns a human friendly text of the log level.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	}
}
