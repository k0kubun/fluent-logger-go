package fluent

import (
	"net"
	"strconv"
)

type Logger struct {
	config Config
	postCh chan Message
	buffer []byte
	conn   net.Conn
}

func NewLogger(config Config) *Logger {
	config.applyDefaultValues()

	logger := &Logger{
		config: config,
		postCh: make(chan Message, config.channelLength),
	}
	go logger.loop()

	return logger
}

func (l *Logger) Post(tag string, data interface{}) {
	l.postCh <- Message{tag: tag, data: data}
}

func (l *Logger) loop() {
	for {
		select {
		case message := <-l.postCh:
			pack, err := message.toMsgpack()
			if err != nil {
				println("error")
				continue
			}

			l.buffer = append(l.buffer, pack...)
			if len(l.buffer) > l.config.bufferLength {
				l.sendMessage()
			}
		}
	}
}

func (l *Logger) sendMessage() {
}

func (l *Logger) connect() {
	for i := 0; i < l.config.maxTrialForConnection; i++ {
		conn, err := net.DialTimeout(
			"tcp",
			l.config.fluentHost+":"+strconv.Itoa(l.config.fluentPort),
			l.config.connectionTimeout,
		)
		if err == nil {
			l.conn = conn
		}
	}
}
