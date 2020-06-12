package codec

type MessageType int

const (
	Request MessageType = iota
	Response
	Event
	Error
)

func (m MessageType) String() string {
	switch m {
	case Request:
		return "Request"
	case Response:
		return "Response"
	case Event:
		return "Event"
	case Error:
		return "Error"
	default:
		return "Unknown"
	}
}

type Message struct {
	Id       string
	Type     MessageType
	Target   string
	Method   string
	Endpoint string
	Error    string
	Header   map[string]string
	Body     []byte
}

type Reader interface {
	ReadHeader(*Message, MessageType) error
	ReadBody(interface{}) error
}

type Writer interface {
	Write(*Message, interface{}) error
}

type Codec interface {
	Reader
	Writer
	Close() error
	String() string
}
