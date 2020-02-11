package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/y-yagi/kura"
)

var (
	logger *kura.Logger
)

func main() {
	logger = kura.NewLogger(os.Stdout)

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
			Name:      "new",
			Aliases:   []string{"n"},
			Usage:     "create a new module",
			ArgsUsage: "[module name]",
			Action:    runNew,
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "lib", Usage: "use a binary template"},
				&cli.BoolFlag{Name: "bin", Usage: "use a library template (default)"},
				&cli.BoolFlag{Name: "no-mod-init", Usage: "do not run 'mod init'"},
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build a module",
			Action:  runBuild,
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "release", Usage: "build in release mode"},
				&cli.StringFlag{Name: "ldflags", Usage: "`ldflags` to specify a build"},
			},
		},
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "install a module",
			Action:  runInstall,
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "release", Usage: "build in release mode"},
				&cli.StringFlag{Name: "ldflags", Usage: "`ldflags` to specify a build"},
			},
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run a module",
			Action:  runRun,
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "release", Usage: "build in release mode"},
				&cli.StringFlag{Name: "ldflags", Usage: "`ldflags` to specify a build"},
			},
		},
	}
}

func appRun(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return nil
}
