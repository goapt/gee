package gee

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	VERSION    = "v1.1.5"
	AppContext context.Context
)

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
	app.Copyright = "(c) Gee"
	app.Writer = os.Stdout
	cli.ErrWriter = os.Stdout

	app.Commands = cmds
	app.Setup()

	ctx, cancel := context.WithCancel(context.Background())
	AppContext = ctx
	var args []string
	args = append(args, app.Name)
	args = append(args, CliArgs...)

	RegisterShutDown(func(sig os.Signal) {
		cancel()
	})

	go WaitSignal()
	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}
