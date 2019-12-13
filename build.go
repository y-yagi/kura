package main

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func runBuild(c *cli.Context) error {
	buildOpt := []string{"build"}
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

	out, err := exec.Command("go", buildOpt...).CombinedOutput()
	if err != nil {
		fmt.Printf("Build: %s\n", out)
		return err
	}
	return nil
}
