package fluent

import (
	"testing"
	"time"
)

func TestApplyDefaultValues(t *testing.T) {
	config := Config{}

	assertEqual(t, config.FluentHost, stringDefault)
	assertEqual(t, config.FluentPort, intDefault)
	assertEqual(t, config.ChannelLength, intDefault)
	assertEqual(t, config.BufferLength, intDefault)
	assertEqual(t, config.MaxTrialForConnection, intDefault)
	assertEqual(t, config.ConnectionTimeout, durationDefault)

	config.applyDefaultValues()
	assertEqual(t, config.FluentHost, defaultFluentHost)
	assertEqual(t, config.FluentPort, defaultFluentPort)
	assertEqual(t, config.ChannelLength, defaultChannelLength)
	assertEqual(t, config.BufferLength, defaultBufferLength)
	assertEqual(t, config.MaxTrialForConnection, defaultMaxTrialForConnection)
	assertEqual(t, config.ConnectionTimeout, defaultConnectionTimeout)

	config = Config{
		FluentHost:            "localhost",
		FluentPort:            80,
		ChannelLength:         1,
		BufferLength:          2,
		MaxTrialForConnection: 3,
		ConnectionTimeout:     2 * time.Second,
	}

	config.applyDefaultValues()
	assertEqual(t, config.FluentHost, "localhost")
	assertEqual(t, config.FluentPort, 80)
	assertEqual(t, config.ChannelLength, 1)
	assertEqual(t, config.BufferLength, 2)
	assertEqual(t, config.MaxTrialForConnection, 3)
	assertEqual(t, config.ConnectionTimeout, 2*time.Second)
}

func assertEqual(t *testing.T, actual interface{}, expect interface{}) {
	switch actual.(type) {
	case string:
		if actual.(string) != expect.(string) {
			t.Errorf("expected %s, but got %s\n", expect.(string), actual.(string))
		}
	case int:
		if actual.(int) != expect.(int) {
			t.Errorf("expected %d, but got %d\n", expect.(int), actual.(int))
		}
	case time.Duration:
		if actual.(time.Duration) != expect.(time.Duration) {
			t.Errorf("expected %d, but got %d\n", expect.(time.Duration), actual.(time.Duration))
		}
	}
}
