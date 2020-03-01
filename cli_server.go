package gee

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	VERSION = "v1.2.1"

	AppContext context.Context
	appCancel  context.CancelFunc
)

func init() {
	AppContext, appCancel = context.WithCancel(context.Background())
}

type CliServer struct {
	serv *cli.App
}

func NewCliServer() *CliServer {
	return &CliServer{
		serv: cli.NewApp(),
	}
}

func (h *CliServer) Serv() *cli.App {
	return h.serv
}

func (h *CliServer) Run(cmds cli.Commands) {

	app := h.serv
	app.Name = "app"
	app.Version = VERSION
	app.Copyright = "(c) VeryStar"
	app.Writer = os.Stdout
	cli.ErrWriter = os.Stdout

	app.Commands = cmds
	var args []string
	args = append(args, app.Name)
	args = append(args, CliArgs...)
	waitCh := make(chan struct{})

	stopSignals := make(chan os.Signal, 1)
	app.After = func(c *cli.Context) error {
		stopSignals <- nil
		return nil
	}

	go WaitSignal(waitCh, stopSignals)
	if err := app.Run(args); err != nil {
		log.Fatal(err)
	}
	<-waitCh
}
