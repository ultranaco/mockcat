package main

import (
	"os"
	"path/filepath"
	"strings"
)

// FileDescriptor File metadata
type FileDescriptor struct {
	Name      string
	FullName  string
	Path      string
	FullPath  string
	Extension string
}

// New build new file descriptor
func (f *FileDescriptor) New(file string) {
	f.buildFullPath(file)
	f.exists()
	f.buildExtension()
	f.buildPath()
	f.buildName()
}

func (f *FileDescriptor) buildFullPath(file string) {
	var currentPath string

	if !filepath.IsAbs(file) {
		path, _ := os.Getwd()
		currentPath = path + string(filepath.Separator)
	}

	f.FullPath = filepath.Clean(currentPath + file)
}

func (f *FileDescriptor) exists() {
	fi, err := os.Stat(f.FullPath)
	if err != nil || fi.IsDir() {
		Error().Raise(1, f.FullPath)
	}
}

func (f *FileDescriptor) buildExtension() {
	f.Extension = filepath.Ext(f.FullPath)
}

func (f *FileDescriptor) buildPath() {
	f.Path = filepath.Dir(f.FullPath)
}

func (f *FileDescriptor) buildName() {
	f.FullName = filepath.Base(f.FullPath)
	f.Name = strings.Replace(f.FullName, f.Extension, "", 1)
}
