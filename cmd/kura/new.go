package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

const bintpl = `package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
`

const libtpl = `package {{.Package}}

func {{.Package}}(name string) string {
	return "Hello, " + name
}
`

const libtesttpl = `package {{.Package}}

import (
	"testing"
)

var tests = []struct {
	in   string
	want string
}{
	{"Bob", "Hello, Bob"},
}

func Test{{.TitleizePackage}}(t *testing.T) {
	for _, tt := range tests {
	  got := {{.Package}}(tt.in)
	  if tt.want != got {
	  	t.Fatalf("got %+v, want %+v", got, tt.want)
	  }
	}
}
`

type content struct {
	Package         string
	TitleizePackage string
}

func runNew(c *cli.Context) error {
	module := c.Args().Get(0)
	if len(module) == 0 {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	a := strings.Split(module, "/")
	packageName := a[len(a)-1]

	if err := os.Mkdir(packageName, 0755); err != nil {
		return err
	}

	var buffer bytes.Buffer
	if c.Bool("lib") {
		content := content{Package: packageName, TitleizePackage: strings.Title(packageName)}

		t := template.Must(template.New("main").Parse(libtpl))
		if err := t.Execute(&buffer, content); err != nil {
			return err
		}
		file := filepath.Join(packageName, packageName+".go")
		if err := ioutil.WriteFile(file, buffer.Bytes(), 0644); err != nil {
			return err
		}

		buffer.Reset()
		t = template.Must(template.New("test").Parse(libtesttpl))
		if err := t.Execute(&buffer, content); err != nil {
			return err
		}
		file = filepath.Join(packageName, packageName+"_test.go")
		if err := ioutil.WriteFile(file, buffer.Bytes(), 0644); err != nil {
			return err
		}
	} else {
		t := template.Must(template.New("main").Parse(bintpl))
		if err := t.Execute(&buffer, nil); err != nil {
			return err
		}

		file := filepath.Join(packageName, "main.go")
		if err := ioutil.WriteFile(file, buffer.Bytes(), 0644); err != nil {
			return err
		}
	}

	os.Chdir(packageName)
	if !c.Bool("no-mod-init") {
		cmd := exec.Command("go", "mod", "init", module)
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	logger.Printf("Created", "'%s' module \n", module)
	return nil
}
