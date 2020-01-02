package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func runBuild(c *cli.Context) error {
	return buildOrInstall("build", c)
}

func runInstall(c *cli.Context) error {
	return buildOrInstall("install", c)
}

func buildOrInstall(action string, c *cli.Context) error {
	buildOpt := []string{}
	ldflags := ""

	if !c.Bool("release") {
		buildOpt = append(buildOpt, "-gcflags")
		buildOpt = append(buildOpt, "-N -l")
	}

	if c.Bool("release") {
		ldflags += "-w -s "
	}

	if len(c.String("ldflags")) != 0 {
		ldflags += c.String("ldflags")
	}

	if len(ldflags) != 0 {
		buildOpt = append(buildOpt, "-ldflags")
		buildOpt = append(buildOpt, ldflags)
	}

	optWithCommand := append([]string{action}, buildOpt...)
	out, err := exec.Command("go", optWithCommand...).CombinedOutput()
	if err != nil {
		fmt.Printf("%s failed: %s\n", strings.Title(action), out)
		return err
	}
	logger.Printf(strings.Title(action)+" Finished", "with '%v' options\n", buildOpt)
	return nil
}
