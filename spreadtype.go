package main

// SpreadType Way to serve content from http handler
type SpreadType int

const (
	// Text type
	Text SpreadType = iota
	// JSON type
	JSON
)
