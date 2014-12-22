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
	DefaultTagPrefix             = ""
)

var (
	intDefault      int
	stringDefault   string
	durationDefault time.Duration
)

// Config is just for fluent.NewLogger() argument.
type Config struct {
	// You can customize fluentd host and port.
	FluentHost string
	FluentPort int

	// If buffered channel's length is equal to ChannelLength, main thread blocks.
	ChannelLength int

	// If all posted messages' size reaches BufferLength, logger flushes all logs.
	BufferLength int

	// Retry connection with fluentd MaxTrialForConnection times.
	MaxTrialForConnection int

	// Wait for connection until ConnectionTimeout.
	ConnectionTimeout time.Duration

	// Logger flushes its buffer on each BufferingTimeout interval.
	BufferingTimeout time.Duration

	// Tag prefix. When set TagPrefix to "foo" and post with "bar.baz",
	// you'll get "foo.bar.baz" tag.
	TagPrefix string
}

func (c *Config) applyDefaultValues() {
	assignIfDefault(&c.FluentHost, DefaultFluentHost)
	assignIfDefault(&c.FluentPort, DefaultFluentPort)
	assignIfDefault(&c.ChannelLength, DefaultChannelLength)
	assignIfDefault(&c.BufferLength, DefaultBufferLength)
	assignIfDefault(&c.MaxTrialForConnection, DefaultMaxTrialForConnection)
	assignIfDefault(&c.ConnectionTimeout, DefaultConnectionTimeout)
	assignIfDefault(&c.BufferingTimeout, DefaultBufferingTimeout)
	assignIfDefault(&c.TagPrefix, DefaultTagPrefix)
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
