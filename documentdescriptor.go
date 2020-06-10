package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"path/filepath"
	"regexp"
)

const (
	// DefaultContentType default mime type
	DefaultContentType = "text/plain"
)

// DocumentDescriptor configuration to make a http handler
type DocumentDescriptor struct {
	Route        string
	ContentType  string
	ResponseBody *interface{}
	IsCollection bool
	FieldMatcher string
	SpreadType   SpreadType
}

//New intialize a request handler descriptor from file descriptor
func (r *DocumentDescriptor) New(descriptor FileDescriptor) {
	r.buildRoute(descriptor)
	r.buildContentType(descriptor)
	r.buildFieldMatcher(descriptor)
	r.buildResponseBody(descriptor)
}

// SetFieldMatcher set field to make searches
func (r *DocumentDescriptor) SetFieldMatcher(fieldMatcher string) {

	cleanPattern := regexp.MustCompile(`[^a-zA-Z0-9_]+`)
	fieldMatcher = cleanPattern.ReplaceAllString(fieldMatcher, "")
	if len(fieldMatcher) == 0 {
		log.Printf("-m '%s' field matcher must have at least one letter or number", fieldMatcher)
	} else {
		r.FieldMatcher = fieldMatcher
	}
}

func (r *DocumentDescriptor) buildRoute(descriptor FileDescriptor) {
	routePattern := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	route := routePattern.ReplaceAllString(descriptor.Name, "/")
	route = filepath.Clean(fmt.Sprintf("/%s", route))

	r.Route = route + "/"
}

func (r *DocumentDescriptor) buildContentType(descriptor FileDescriptor) {
	r.ContentType = DefaultContentType
	mimeType := mime.TypeByExtension(descriptor.Extension)
	if len(mimeType) > 0 {
		r.ContentType = mimeType
	}
}

func (r *DocumentDescriptor) buildFieldMatcher(descriptor FileDescriptor) {
	r.FieldMatcher = "id"
}

func (r *DocumentDescriptor) buildResponseBody(descriptor FileDescriptor) {
	var j interface{}
	content, err := ioutil.ReadFile(descriptor.FullPath)

	if err != nil {
		Error().Raise(2, descriptor.FullPath)
	}

	_ = json.Unmarshal(content, &j)
	_, mOk := j.(map[string]interface{})
	_, sOk := j.([]interface{})

	r.ResponseBody = &j

	if mOk {
		r.IsCollection = false
		r.SpreadType = JSON
	} else if sOk {
		r.IsCollection = true
		r.SpreadType = JSON
	} else {
		r.SpreadType = Text
		var cnt interface{} = string(content)
		r.ResponseBody = &cnt
	}
}
