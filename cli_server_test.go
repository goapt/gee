package gee

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewCliServer(t *testing.T) {
	s := NewCliServer()
	assert.Equal(t, s.serv, s.Serv())
}

func TestCliServer_Run(t *testing.T) {
	s := NewCliServer()
	// cmd handler
	var testCmd = &cli.Command{
		Name:  "test",
		Usage: "test command eg: ./app test --id=7",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "id",
				Usage: "user id",
				Value: "7",
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println(ctx.Int("id"))
			return nil
		},
	}

	assert.NotPanics(t, func() {
		s.Run([]*cli.Command{testCmd})
	})
}
