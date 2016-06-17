package fluent

import (
	"errors"
	"log"
	"net"
	"strconv"
	"time"
)

// Logger owns asynchronous logging to fluentd.
type Logger struct {
	config   Config
	postCh   chan message
	buffer   []byte
	conn     net.Conn
	ticker   *time.Ticker
	logError bool
}

// NewLogger() launches a goroutine to log and returns logger.
// Logger has a channel to interact with the goroutine.
func NewLogger(config Config) *Logger {
	config.applyDefaultValues()

	logger := &Logger{
		config:   config,
		postCh:   make(chan message, config.ChannelLength),
		buffer:   make([]byte, 0, config.BufferLength),
		ticker:   time.NewTicker(config.BufferingTimeout),
		logError: true,
	}
	logger.connect()
	go logger.loop()

	return logger
}

// You can send message to logger's goroutine via channel.
// This logging is executed asynchronously.
func (l *Logger) Post(tag string, data interface{}) {
	tag = l.prependTagPrefix(tag)
	l.postCh <- message{tag: tag, time: time.Now(), data: data}
}

// You can send message immediately to fluentd.
func (l *Logger) Log(tag string, data interface{}) error {
	tag = l.prependTagPrefix(tag)
	msg := &message{tag: tag, time: time.Now(), data: data}
	pack, err := msg.toMsgpack()
	if err != nil {
		return err
	}

	l.buffer = append(l.buffer, pack...)
	return l.sendMessage()
}

func (l *Logger) loop() {
	for {
		select {
		case msg := <-l.postCh:
			pack, err := msg.toMsgpack()
			if err != nil {
				log.Printf("message pack dump error: " + err.Error())
				continue
			}

			l.buffer = append(l.buffer, pack...)
			if len(l.buffer) > l.config.BufferLength {
				l.sendMessage()
			}
		case <-l.ticker.C:
			l.sendMessage()
		}
	}
}

func (l *Logger) sendMessage() error {
	if len(l.buffer) == 0 {
		return errors.New("Buffer is empty")
	}

	l.connect()
	if l.conn == nil {
		return errors.New("Failed to establish connection with fluentd")
	}

	_, err := l.conn.Write(l.buffer)

	if err == nil {
		l.buffer = l.buffer[0:0]
	} else {
		log.Printf("failed to send message: " + err.Error())
		l.conn.Close()
		l.conn = nil
	}
	return err
}

func (l *Logger) connect() {
	if l.conn != nil {
		return
	}

	var err error
	for i := 0; i < l.config.MaxTrialForConnection; i++ {
		l.conn, err = net.DialTimeout(
			"tcp",
			l.config.FluentHost+":"+strconv.Itoa(l.config.FluentPort),
			l.config.ConnectionTimeout,
		)

		if err == nil {
			l.logError = true
			return
		}
	}

	if l.logError {
		log.Printf("failed to establish connection with fluentd: " + err.Error())
		l.logError = false
	}
}

func (l *Logger) prependTagPrefix(tag string) string {
	if l.config.TagPrefix != "" {
		tag = l.config.TagPrefix + "." + tag
	}
	return tag
}
