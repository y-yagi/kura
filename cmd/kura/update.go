package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

func runUpdate(c *cli.Context) error {
	out, err := exec.Command("go", "get", "-u", "./...").CombinedOutput()
	if err != nil {
		fmt.Printf("%s failed: %s\n", strings.Title("update"), out)
		return err
	}

	logger.Printf("", string(out))
	return nil
}
