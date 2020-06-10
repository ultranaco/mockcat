package main

import (
	"fmt"
	"net/http"
)

// DocumentHandleText handle incoming request by http Method
type DocumentHandleText struct {
	Document DocumentDescriptor
}

// New initialize
func (d *DocumentHandleText) New(document DocumentDescriptor) {
	d.Document = document
}

//Get GET verb
func (d *DocumentHandleText) Get(key string) HTTPResponse {
	response := fmt.Sprintf("%v", *d.Document.ResponseBody)
	return HTTPResponse{
		Status: http.StatusOK,
		Body:   response,
	}
}

//Delete DELETE verb
func (d *DocumentHandleText) Delete(key string) HTTPResponse {
	response := "DELETE method not allowed with non iterable data"
	return HTTPResponse{
		Status: http.StatusMethodNotAllowed,
		Body:   response,
	}
}

// Post POST verb
func (d *DocumentHandleText) Post(payload string) HTTPResponse {
	response := "POST method not allowed with non iterable data"
	return HTTPResponse{
		Status: http.StatusMethodNotAllowed,
		Body:   response,
	}
}

//Put PUT verb
func (d *DocumentHandleText) Put(key string, payload string) HTTPResponse {
	response := "PUT method not allowed with non iterable data"
	return HTTPResponse{
		Status: http.StatusMethodNotAllowed,
		Body:   response,
	}
}
