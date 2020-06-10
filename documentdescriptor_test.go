package main

import (
	"testing"
)

func TestNewRQHandle(t *testing.T) {
	fd := FileDescriptor{}
	reqDescriptor := DocumentDescriptor{}

	fd.New("testdata/food-mexican(url).json")
	reqDescriptor.New(fd)
}

func TestNewTextFile(t *testing.T) {
	fd := FileDescriptor{}
	reqDescriptor := DocumentDescriptor{}

	fd.New("testdata/food-mexican.txt")
	reqDescriptor.New(fd)
}

func TestNewRQNoFormatFile(t *testing.T) {
	fd := FileDescriptor{}
	reqDescriptor := DocumentDescriptor{}

	fd.New("testdata/food +mexican_url.bar.json")
	reqDescriptor.New(fd)
}

func TestNewRQOnlyObject(t *testing.T) {
	fd := FileDescriptor{}
	reqDescriptor := DocumentDescriptor{}

	fd.New("testdata/food-mexican-item.json")
	reqDescriptor.New(fd)
}

func TestNewRQNoExtension(t *testing.T) {
	fd := FileDescriptor{}
	reqDescriptor := DocumentDescriptor{}

	fd.New("testdata/food-mexican(url)")
	reqDescriptor.New(fd)
}
