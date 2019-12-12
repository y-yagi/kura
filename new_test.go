package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestNew_Bin(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "newtest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	os.Chdir(tempDir)
	run([]string{"kura", "new", "-m", "github.com/y-yagi/dummy"})
	os.Chdir("dummy")

	got, err := exec.Command("go", "run", "main.go").Output()
	if err != nil {
		log.Fatal(err)
	}

	want := "Hello, world!\n"
	if string(got) != want {
		t.Fatalf("got '%q', want '%q'", got, want)
	}
}

func TestNew_Lib(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "newtest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	os.Chdir(tempDir)
	run([]string{"kura", "new", "--lib", "-m", "github.com/y-yagi/dummy"})
	os.Chdir("dummy")

	got, err := exec.Command("go", "test").Output()
	if err != nil {
		log.Fatal(err)
	}

	want := "PASS\nok"
	if !strings.HasPrefix(string(got), want) {
		t.Fatalf("got '%q', want '%q'", got, want)
	}
}
