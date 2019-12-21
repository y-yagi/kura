package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/y-yagi/goext/ioext"
)

func TestBuild(t *testing.T) {
	tempDir, err := setupTestFile()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	runBuildWithCleanEnv(tempDir, []string{})

	out, err := exec.Command("file", filepath.Join(tempDir, filepath.Base(tempDir))).CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(out), "not stripped") {
		t.Fatalf("expect '%q' does include 'not stripped'", out)
	}
}

func TestBuild_Release(t *testing.T) {
	tempDir, err := setupTestFile()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	runBuildWithCleanEnv(tempDir, []string{"--release"})

	out, err := exec.Command("file", filepath.Join(tempDir, filepath.Base(tempDir))).CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	if strings.Contains(string(out), "not stripped") {
		t.Fatalf("expect '%q' does not include 'not stripped'", out)
	}
}

func TestBuild_Ldflags(t *testing.T) {
	tempDir, err := setupTestFile()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	runBuildWithCleanEnv(tempDir, []string{"--ldflags", "-X main.version=v1.0.0"})

	out, err := exec.Command(filepath.Join(tempDir, filepath.Base(tempDir))).CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	want := "version: v1.0.0"
	if string(out) != want {
		t.Fatalf("expect '%q' but got '%q'", want, string(out))
	}
}

func runBuildWithCleanEnv(path string, opts []string) {
	cmd := []string{"kura", "build"}
	if len(opts) != 0 {
		cmd = append(cmd, opts...)
	}

	origEnv := os.Getenv("GO111MODULE")
	origPath, _ := filepath.Abs(".")

	os.Setenv("GO111MODULE", "off")
	os.Chdir(path)

	run(cmd)

	os.Setenv("GO111MODULE", origEnv)
	os.Chdir(origPath)
}

func setupTestFile() (string, error) {
	tempDir, err := ioutil.TempDir("", "buildtest")
	if err != nil {
		return "", err
	}

	absPath := "../../"
	if err = ioext.CopyFile(filepath.Join(absPath, "testdata", "build_test.go"), filepath.Join(tempDir, "main.go"), 0644); err != nil {
		return "", err
	}

	return tempDir, nil
}
