package protobuf

import (
	"errors"
	"github.com/gogo/protobuf/proto"
	"github.com/xdatk/micro/codec"
	"io"
	"io/ioutil"
)

var ErrorNotProtobufStruct = errors.New("not protobuf struct")

type Codec struct {
	Conn io.ReadWriteCloser
}

func (c *Codec) ReadHeader(*codec.Message, codec.MessageType) error {
	return nil
}

func (c *Codec) ReadBody(b interface{}) error {
	if b == nil {
		return nil
	}

	m, ok := b.(proto.Message)
	if !ok {
		return ErrorNotProtobufStruct
	}

	buf, err := ioutil.ReadAll(c.Conn)
	if err != nil {
		return err
	}

	return proto.Unmarshal(buf, m)
}

func (c *Codec) Write(_ *codec.Message, b interface{}) error {
	p, ok := b.(proto.Message)
	if !ok {
		return ErrorNotProtobufStruct
	}
	buf, err := proto.Marshal(p)
	if err != nil {
		return err
	}

	_, err = c.Conn.Write(buf)
	return err
}

func (c *Codec) Close() error {
	return c.Conn.Close()
}

func (c *Codec) String() string {
	return "protobuf"
}

func NewCodec(c io.ReadWriteCloser) codec.Codec {
	return &Codec{
		Conn: c,
	}
}
