package very

import (
	"context"
	"github.com/urfave/cli"
)

var Commands []cli.Command

func Register(cmd cli.Command) {
	Commands = append(Commands, cmd)
}

func GetContext(c *cli.Context) context.Context {
	return c.App.Metadata["ctx"].(context.Context)
}
