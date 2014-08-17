package fluent

import (
	"testing"
	"time"
)

func TestApplyDefaultValues(t *testing.T) {
	config := Config{}

	assertEqual(t, config.fluentHost, stringDefault)
	assertEqual(t, config.fluentPort, intDefault)
	assertEqual(t, config.channelLength, intDefault)
	assertEqual(t, config.bufferLength, intDefault)
	assertEqual(t, config.maxTrialForConnection, intDefault)
	assertEqual(t, config.connectionTimeout, durationDefault)

	config.applyDefaultValues()
	assertEqual(t, config.fluentHost, defaultFluentHost)
	assertEqual(t, config.fluentPort, defaultFluentPort)
	assertEqual(t, config.channelLength, defaultChannelLength)
	assertEqual(t, config.bufferLength, defaultBufferLength)
	assertEqual(t, config.maxTrialForConnection, defaultMaxTrialForConnection)
	assertEqual(t, config.connectionTimeout, defaultConnectionTimeout)

	config = Config{
		fluentHost:            "localhost",
		fluentPort:            80,
		channelLength:         1,
		bufferLength:          2,
		maxTrialForConnection: 3,
		connectionTimeout:     2 * time.Second,
	}

	config.applyDefaultValues()
	assertEqual(t, config.fluentHost, "localhost")
	assertEqual(t, config.fluentPort, 80)
	assertEqual(t, config.channelLength, 1)
	assertEqual(t, config.bufferLength, 2)
	assertEqual(t, config.maxTrialForConnection, 3)
	assertEqual(t, config.connectionTimeout, 2*time.Second)
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
