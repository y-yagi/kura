package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/y-yagi/goext/ioext"
	"github.com/y-yagi/kura"
)

var b bytes.Buffer

func TestMain(m *testing.M) {
	logger = kura.NewLogger(&b)
	os.Exit(m.Run())
}

func TestBuild(t *testing.T) {
	tempDir, err := setupTestFile()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	runCommandWithCleanEnv("build", tempDir, []string{})

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

	runCommandWithCleanEnv("build", tempDir, []string{"--release"})

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

	runCommandWithCleanEnv("build", tempDir, []string{"--ldflags", "-X main.version=v1.0.0"})

	out, err := exec.Command(filepath.Join(tempDir, filepath.Base(tempDir))).CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	want := "version: v1.0.0"
	if string(out) != want {
		t.Fatalf("expect '%q' but got '%q'", want, string(out))
	}
}

func TestInstall(t *testing.T) {
	tempDir, err := setupTestFile()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	runCommandWithCleanEnv("install", tempDir, []string{})

	out, err := exec.Command("file", filepath.Join(tempDir, filepath.Base(tempDir))).CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(out), "not stripped") {
		t.Fatalf("expect '%q' does include 'not stripped'", out)
	}
}

func TestInstall_Release(t *testing.T) {
	tempDir, err := setupTestFile()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	runCommandWithCleanEnv("install", tempDir, []string{"--release"})

	out, err := exec.Command("file", filepath.Join(tempDir, filepath.Base(tempDir))).CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	if strings.Contains(string(out), "not stripped") {
		t.Fatalf("expect '%q' does not include 'not stripped'", out)
	}
}

func TestInstall_Ldflags(t *testing.T) {
	tempDir, err := setupTestFile()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	runCommandWithCleanEnv("install", tempDir, []string{"--ldflags", "-X main.version=v1.0.0"})

	out, err := exec.Command(filepath.Join(tempDir, filepath.Base(tempDir))).CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	want := "version: v1.0.0"
	if string(out) != want {
		t.Fatalf("expect '%q' but got '%q'", want, string(out))
	}
}

func TestRun(t *testing.T) {
	tempDir, err := setupTestFile()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	runCommandWithCleanEnv("run", tempDir, []string{"."})

	out := b.String()
	if !strings.Contains(out, "version: dev") {
		t.Fatalf("expect '%q' does include 'version: dev'", out)
	}
}

func runCommandWithCleanEnv(action, path string, opts []string) {
	cmd := []string{"kura", action}
	if len(opts) != 0 {
		cmd = append(cmd, opts...)
	}

	origModule := os.Getenv("GO111MODULE")
	origBin := os.Getenv("GOBIN")
	origPath, _ := filepath.Abs(".")

	os.Setenv("GO111MODULE", "off")
	os.Setenv("GOBIN", path)
	os.Chdir(path)

	run(cmd)

	os.Setenv("GO111MODULE", origModule)
	os.Setenv("GOBIN", origBin)
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
