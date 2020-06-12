package cmd

import (
	"github.com/urfave/cli/v2"
	"os"
)

var (
	DefaultFlags = []cli.Flag{}
)

type Cmd interface {
	App() *cli.App
	Init(...Option) error
	Options() Options
}

type Option func(o *Options)

func NewCmd(opts ...Option) Cmd {
	options := Options{}

	for _, opt := range opts {
		opt(&options)
	}

	if len(options.Description) == 0 {
		options.Description = "a micro service"
	}

	cmd := new(cmd)
	cmd.options = options
	cmd.app = cli.NewApp()
	cmd.app.Name = cmd.options.Name
	cmd.app.Version = cmd.options.Version
	cmd.app.Usage = cmd.options.Usage
	cmd.app.Before = cmd.Before
	cmd.app.Flags = DefaultFlags
	cmd.app.Action = func(c *cli.Context) error {
		return nil
	}

	if len(options.Version) == 0 {
		cmd.app.HideVersion = true
	}

	return cmd
}

type cmd struct {
	options Options
	app     *cli.App
}

func (c *cmd) App() *cli.App {
	return c.app
}

func (c *cmd) Init(opts ...Option) error {
	for _, opt := range opts {
		opt(&c.options)
	}

	if len(c.options.Name) > 0 {
		c.app.Name = c.options.Name
	}

	if len(c.options.Version) > 0 {
		c.app.Version = c.options.Version
	}

	c.app.HideVersion = len(c.options.Version) == 0
	c.app.Usage = c.options.Description
	return c.app.Run(os.Args)
}

func (c *cmd) Options() Options {
	return c.options
}

func (c *cmd) Before(ctx *cli.Context) error {
	return nil
}
