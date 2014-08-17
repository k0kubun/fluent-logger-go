package fluent

import (
	"time"

	"github.com/ugorji/go/codec"
)

var (
	mh codec.MsgpackHandle
)

type message struct {
	tag  string
	time time.Time
	data interface{}
}

func (m *message) toMsgpack() ([]byte, error) {
	pack := []byte{}
	encoder := codec.NewEncoderBytes(&pack, &mh)

	rawMessage := []interface{}{m.tag, m.time.Unix(), m.data}
	err := encoder.Encode(rawMessage)

	return pack, err
}
