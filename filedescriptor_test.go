package main

import (
	"os"
	"os/exec"
	"testing"
)

type expectExitError func()

func TestNewFileDescriptor(t *testing.T) {
	fileDescriptor := FileDescriptor{}
	fileDescriptor.New("../mockcat/../mockcat/testdata/food-mexican(url).json")
}

func TestNewFileNotExists(t *testing.T) {
	fileDescriptor := FileDescriptor{}
	expectExit(func() {
		fileDescriptor.New("testdata/noexists.json")
	}, "TestNewFileWitDirectory", t)
}

func TestNewFileWithoutExtension(t *testing.T) {
	fileDescriptor := FileDescriptor{}
	fileDescriptor.New("testdata/food-mexican(url)")
}

func TestNewFileWitDirectory(t *testing.T) {
	fileDescriptor := FileDescriptor{}

	expectExit(func() {
		fileDescriptor.New("testdata/food-mexican_url-dir")
	}, "TestNewFileWitDirectory", t)
}

func expectExit(fn expectExitError, testName string, t *testing.T) {
	if os.Getenv("EXIT") == "1" {
		fn()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run="+testName)
	cmd.Env = append(os.Environ(), "EXIT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("Process ran with err %v, want exit status 1", err)
}
