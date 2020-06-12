package micro

import (
	"github.com/xdatk/micro/client"
	"github.com/xdatk/micro/server"
	signalUtil "github.com/xdatk/micro/util/signal"
	"os"
	"os/signal"
	"sync"
)

type service struct {
	options Options
	once    sync.Once
}

func (s *service) Name() string {
	return ""
}

func (s *service) Init(opts ...Option) {
	for _, opt := range opts {
		opt(&s.options)
	}

	s.once.Do(func() {
		if len(s.options.Cmd.App().Name) == 0 {
			s.options.Cmd.App().Name = s.options.Server.Options().Name
		}

		err := s.options.Cmd.Init()
		if err != nil {
			panic(err)
		}
	})
}

func (s *service) Options() Options {
	return s.options
}

func (s *service) Client() client.Client {
	return s.options.Client
}

func (s *service) Server() server.Server {
	return s.options.Server
}

func (s *service) Run() error {

	if err := s.Start(); err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	if s.options.Signal {
		signal.Notify(ch, signalUtil.Shutdown()...)
	}

	select {
	case <-ch:
	case <-s.options.Context.Done():
	}

	return s.Stop()
}

func (s *service) Start() error {
	for _, fn := range s.options.BeforeStart {
		if err := fn(); err != nil {
			return err
		}
	}

	if err := s.options.Server.Start(); err != nil {
		return err
	}

	for _, fn := range s.options.AfterStart {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func (s *service) Stop() (err error) {

	for _, fn := range s.options.BeforeStop {
		err = fn()
	}

	err = s.options.Server.Stop()

	for _, fn := range s.options.AfterStop {
		err = fn()
	}

	return
}

func (s *service) String() string {
	return "micro"
}
