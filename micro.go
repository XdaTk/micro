package micro

import (
	"github.com/xdatk/micro/client"
	"github.com/xdatk/micro/server"
)

type Service interface {
	Name() string
	Init(...Option)
	Options() Options
	Client() client.Client
	Server() server.Server
	Run() error
	String() string
}
