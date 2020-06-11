package micro

import (
	"github.com/xdatk/micro/client"
	"github.com/xdatk/micro/server"
)

// Service is an interface that wraps the lower level libraries
// within micro. Its a convenience method for building
// and initialising services.
type Service interface {
	// The service name
	Name() string
	// Init initialises options
	Init(...Option)
	// Options returns the current options
	Options() Options
	// Client is used to call services
	Client() client.Client
	// Server is for handling requests and events
	Server() server.Server
	// Run the service
	Run() error
	// The service implementation
	String() string
}
