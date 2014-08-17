package fluent

import (
	"log"
	"net"
	"strconv"
	"time"
)

// Logger owns asynchronous logging to fluentd.
type Logger struct {
	config Config
	postCh chan message
	buffer []byte
	conn   net.Conn
	ticker *time.Ticker
}

func NewLogger(config Config) *Logger {
	config.applyDefaultValues()

	logger := &Logger{
		config: config,
		postCh: make(chan message, config.ChannelLength),
		ticker: time.NewTicker(config.BufferingTimeout),
	}
	go logger.loop()
	defer logger.sendMessage()

	return logger
}

func (l *Logger) Post(tag string, data interface{}) {
	l.postCh <- message{tag: tag, data: data}
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

func (l *Logger) sendMessage() {
	if len(l.buffer) == 0 {
		return
	}

	l.connect()
	_, err := l.conn.Write(l.buffer)

	if err == nil {
		l.buffer = l.buffer[0:0]
	} else {
		log.Printf("failed to send message: " + err.Error())
		l.conn.Close()
		l.conn = nil
	}
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
			return
		}
	}
	log.Printf("failed to establish connection with fluentd: " + err.Error())
}
