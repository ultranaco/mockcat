package main

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func initJSON(filePath string) *httptest.Server {
	file := FileDescriptor{}
	doc := DocumentDescriptor{}
	handle := ResponseHandle{}

	file.New(filePath)
	doc.New(file)
	handle.New(doc)

	ts := httptest.NewServer(http.HandlerFunc(handle.Handle))
	return ts
}

func TestGetJSONHandle(t *testing.T) {
	ts := initJSON("testdata/food-mexican(url).json")
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/food/mexican/mole")

	greeting, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("%s", greeting)
}

func TestPostJSONHandle(t *testing.T) {
	ts := initJSON("testdata/food-mexican(url).json")
	defer ts.Close()

	body := strings.NewReader(`{
		"description": "This is the piquito de Gallo",
		"id": 4,
		"title": "Piquito de Gallo very tasty",
		"url": "gallote"
	}`)

	res, _ := http.Post(ts.URL+"/food/mexican/", mime.TypeByExtension(".json"), body)

	greeting, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("%s", greeting)
}

func TestPutJSONHandle(t *testing.T) {
	ts := initJSON("testdata/food-mexican(url).json")
	defer ts.Close()

	body := strings.NewReader(`{
		"description": "This is the piquito de Gallo",
		"id": 4,
		"title": "Piquito de Gallo very tasty",
		"url": "gallote"
	}`)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/food/mexican/mole", body)
	res, _ := client.Do(req)

	greeting, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("%s", greeting)
}
func TestDeleteJSONHandle(t *testing.T) {
	ts := initJSON("testdata/food-mexican(url).json")
	defer ts.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, ts.URL+"/food/mexican/mole", nil)
	res, _ := client.Do(req)

	greeting, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("%s", greeting)
}
