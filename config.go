package fluent

import "time"

const (
	defaultFluentHost            = "127.0.0.1"
	defaultFluentPort            = 24224
	defaultChannelLength         = 1000
	defaultBufferLength          = 10 * 1024
	defaultMaxTrialForConnection = 10
	defaultConnectionTimeout     = time.Second
	defaultBufferingTimeout      = time.Second
)

var (
	intDefault      int
	stringDefault   string
	durationDefault time.Duration
)

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
	assignIfDefault(&c.FluentHost, defaultFluentHost)
	assignIfDefault(&c.FluentPort, defaultFluentPort)
	assignIfDefault(&c.ChannelLength, defaultChannelLength)
	assignIfDefault(&c.BufferLength, defaultBufferLength)
	assignIfDefault(&c.MaxTrialForConnection, defaultMaxTrialForConnection)
	assignIfDefault(&c.ConnectionTimeout, defaultConnectionTimeout)
	assignIfDefault(&c.BufferingTimeout, defaultBufferingTimeout)
}

func assignIfDefault(target interface{}, defaultValue interface{}) {
	switch target.(type) {
	case *string:
		ptr := target.(*string)
		if *ptr == stringDefault {
			*ptr = defaultValue.(string)
		}
	case *int:
		ptr := target.(*int)
		if *ptr == intDefault {
			*ptr = defaultValue.(int)
		}
	case *time.Duration:
		ptr := target.(*time.Duration)
		if *ptr == durationDefault {
			*ptr = defaultValue.(time.Duration)
		}
	}
}
