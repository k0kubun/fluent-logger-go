package fluent

import "time"

const (
	DefaultFluentHost            = "127.0.0.1"
	DefaultFluentPort            = 24224
	DefaultChannelLength         = 1000
	DefaultBufferLength          = 10 * 1024
	DefaultMaxTrialForConnection = 10
	DefaultConnectionTimeout     = 3 * time.Second
	DefaultBufferingTimeout      = 100 * time.Millisecond
)

var (
	intDefault      int
	stringDefault   string
	durationDefault time.Duration
)

// Config is just for fluent.NewLogger() argument.
type Config struct {
	FluentHost            string
	FluentPort            int
	ChannelLength         int
	BufferLength          int
	MaxTrialForConnection int
	ConnectionTimeout     time.Duration
	BufferingTimeout      time.Duration
}

func (c *Config) applyDefaultValues() {
	assignIfDefault(&c.FluentHost, DefaultFluentHost)
	assignIfDefault(&c.FluentPort, DefaultFluentPort)
	assignIfDefault(&c.ChannelLength, DefaultChannelLength)
	assignIfDefault(&c.BufferLength, DefaultBufferLength)
	assignIfDefault(&c.MaxTrialForConnection, DefaultMaxTrialForConnection)
	assignIfDefault(&c.ConnectionTimeout, DefaultConnectionTimeout)
	assignIfDefault(&c.BufferingTimeout, DefaultBufferingTimeout)
}

func assignIfDefault(target interface{}, DefaultValue interface{}) {
	switch target.(type) {
	case *string:
		ptr := target.(*string)
		if *ptr == stringDefault {
			*ptr = DefaultValue.(string)
		}
	case *int:
		ptr := target.(*int)
		if *ptr == intDefault {
			*ptr = DefaultValue.(int)
		}
	case *time.Duration:
		ptr := target.(*time.Duration)
		if *ptr == durationDefault {
			*ptr = DefaultValue.(time.Duration)
		}
	}
}
