package main

// DocumentHandler crud functions
type DocumentHandler interface {
	New(doc DocumentDescriptor)
	Get(key string) HTTPResponse
	Delete(key string) HTTPResponse
	Post(payload string) HTTPResponse
	Put(key string, payload string) HTTPResponse
}
