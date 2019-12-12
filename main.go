package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	os.Exit(run(os.Args))
}

func msg(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		return 1
	}
	return 0
}

func run(args []string) int {
	app := cli.NewApp()
	app.Name = "kura"
	app.Usage = "Module helper for Go"
	app.Version = "0.0.1"
	app.Action = appRun
	app.Commands = commands()

	return msg(app.Run(args))
}

func commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create a new module",
			Action:  createModule,
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "lib", Usage: "use a binary template"},
				&cli.BoolFlag{Name: "bin", Usage: "use a library template (default)"},
				&cli.StringFlag{Name: "module", Usage: "`module` name to use", Aliases: []string{"m"}, Required: true},
			},
		},
	}
}

func appRun(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return nil
}
