package registry

type Type uint8

const (
	Srv Type = iota
	Event
)

func (t Type) String() string {
	switch t {
	case Srv:
		return "Srv"
	case Event:
		return "Event"
	default:
		return "Unknown"
	}
}

type Service struct {
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Endpoints []*Endpoint       `json:"endpoints"`
	Nodes     []*Node           `json:"nodes"`
}

type Node struct {
	Id       string            `json:"id"`
	Address  string            `json:"address"`
	Metadata map[string]string `json:"metadata"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json:"values"`
}

type Endpoint struct {
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Request  *Value            `json:"request"`
	Response *Value            `json:"response"`
	Metadata map[string]string `json:"metadata"`
}
