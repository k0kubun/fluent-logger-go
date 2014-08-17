package fluent

type Config struct {
}

type Logger struct {
	config Config
	buffer []byte
}

func NewLogger(config Config) *Logger {
	return &Logger{
		config: config,
	}
}

func (l *Logger) Post(tag string, message interface{}) {
	print("hello")
}
