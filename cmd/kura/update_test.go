package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/y-yagi/goext/ioext"
)

func TestUpdate(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "updatetest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	absPath, _ := filepath.Abs("../../")
	if err = ioext.CopyFile(filepath.Join(absPath, "testdata", "update_test.go"), filepath.Join(tempDir, "main.go"), 0644); err != nil {
		t.Fatal(err)
	}
	if err = ioext.CopyFile(filepath.Join(absPath, "testdata", "update_test.mod"), filepath.Join(tempDir, "go.mod"), 0644); err != nil {
		t.Fatal(err)
	}

	origModule := os.Getenv("GO111MODULE")
	os.Setenv("GO111MODULE", "on")
	defer os.Setenv("GO111MODULE", origModule)
	ret := run([]string{"kura", "update"})
	if ret != 0 {
		t.Fatal("update failed", ret)
	}
}
