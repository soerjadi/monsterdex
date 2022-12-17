package log

type Level int

type Config struct {
	Level       Level
	LogFilePath string
	Debug       bool
}

// list of log level
const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type KV = map[string]interface{}
