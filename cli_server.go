package gee

import (
	"context"
	"log"
	"os"

	"github.com/goapt/golib/osutil"
	"github.com/urfave/cli"
)

var (
	VERSION = "v1.1.1"
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

func (h *CliServer) Run() {
	app := h.serv
	app.Name = "app"
	app.Version = VERSION
	app.Copyright = "(c) Gee"
	app.Writer = os.Stdout
	cli.ErrWriter = os.Stdout

	app.Commands = Commands
	app.Setup()

	ctx, cancel := context.WithCancel(context.Background())
	app.Metadata["ctx"] = ctx
	var args []string
	args = append(args, app.Name)
	args = append(args, CliArgs...)

	osutil.RegisterShutDown(func(sig os.Signal) {
		cancel()
	})

	go osutil.WaitSignal()
	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}
