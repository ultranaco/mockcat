package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// ResponseHandle receive all incoming request
type ResponseHandle struct {
	Document        DocumentDescriptor
	DocumentHandler DocumentHandler
}

// New initialize response handle
func (r *ResponseHandle) New(descriptor DocumentDescriptor) {
	r.Document = descriptor
	r.DocumentHandler = r.getType(r.Document.SpreadType)
	r.DocumentHandler.New(descriptor)
}

// Handle manage all incoming requests
func (r *ResponseHandle) Handle(w http.ResponseWriter, rq *http.Request) {
	re := regexp.MustCompile(`/.*$`)
	var response HTTPResponse
	var payload string

	pathValue := strings.Replace(rq.URL.Path, r.Document.Route, "", 1)
	pathValue = re.ReplaceAllString(pathValue, "")

	if b, err := ioutil.ReadAll(rq.Body); err == nil {
		payload = string(b)
	}

	switch rq.Method {
	case "GET":
		response = r.DocumentHandler.Get(pathValue)
	case "POST":
		response = r.DocumentHandler.Post(payload)
	case "PUT":
		response = r.DocumentHandler.Put(pathValue, payload)
	case "DELETE":
		response = r.DocumentHandler.Delete(pathValue)
	}

	w.Header().Set("Content-Type", r.Document.ContentType)
	w.WriteHeader(response.Status)
	w.Write([]byte(response.Body))
}

func (r *ResponseHandle) getType(typ SpreadType) DocumentHandler {

	switch typ {
	case JSON:
		return &DocumentHandleJSON{}
	default:
		return &DocumentHandleText{}
	}
}
