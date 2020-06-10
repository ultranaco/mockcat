package main

import (
	"log"
	"os"
)

var errors map[int]ErrorItem = map[int]ErrorItem{
	1: ErrorItem{1, "Error: file does not exists"},
	2: ErrorItem{2, "Error: can not read file"},
}

// ErrorHandle manage errors by codes
type ErrorHandle struct {
}

// Raise an error with parameters if need it
func (e ErrorHandle) Raise(code int, params ...string) {
	log.Println(errors[code].Message, params)
	os.Exit(code)
}

// Error Factory
func Error() *ErrorHandle {
	return &ErrorHandle{}
}
