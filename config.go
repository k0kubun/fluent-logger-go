package fluent

import "time"

const (
	defaultFluentHost            = "127.0.0.1"
	defaultFluentPort            = 24224
	defaultChannelLength         = 1000
	defaultBufferLength          = 1024 * 1024
	defaultMaxTrialForConnection = 10
	defaultConnectionTimeout     = time.Second
)

var (
	intDefault      int
	stringDefault   string
	durationDefault time.Duration
)

type Config struct {
	fluentHost            string
	fluentPort            int
	channelLength         int
	bufferLength          int
	maxTrialForConnection int
	connectionTimeout     time.Duration
}

func (c *Config) applyDefaultValues() {
	assignIfDefault(&c.fluentHost, defaultFluentHost)
	assignIfDefault(&c.fluentPort, defaultFluentPort)
	assignIfDefault(&c.channelLength, defaultChannelLength)
	assignIfDefault(&c.bufferLength, defaultBufferLength)
	assignIfDefault(&c.maxTrialForConnection, defaultMaxTrialForConnection)
	assignIfDefault(&c.connectionTimeout, defaultConnectionTimeout)
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
