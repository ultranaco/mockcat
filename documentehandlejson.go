package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// DocumentHandleJSON handle incoming request by http Method
type DocumentHandleJSON struct {
	Document DocumentDescriptor
}

// New initialize
func (d *DocumentHandleJSON) New(document DocumentDescriptor) {
	d.Document = document
}

//Get GET verb
func (d *DocumentHandleJSON) Get(key string) HTTPResponse {
	status := http.StatusOK
	var response interface{}
	var index int
	response = *d.Document.ResponseBody

	if d.Document.IsCollection {

		if len(key) > 0 {
			index, response = d.search(key)
			if index < 0 {
				status = http.StatusNotFound
			}
		}

		body, _ := d.toString(response)
		return HTTPResponse{
			Status: status,
			Body:   body,
		}
	}

	body, _ := d.toString(response)
	return HTTPResponse{
		Status: status,
		Body:   body,
	}

}

//Delete DELETE verb
func (d *DocumentHandleJSON) Delete(key string) HTTPResponse {
	status := http.StatusMethodNotAllowed
	response := "DELETE method not allowed json document is not a collection"

	if d.Document.IsCollection && key != "" {
		index, _ := d.search(key)
		if index > -1 {
			d.remove(index)
			response = "success"
			status = http.StatusOK
		} else {
			response = "Item Not Found"
			status = http.StatusNotFound
		}
	}

	return HTTPResponse{
		Status: status,
		Body:   response,
	}
}

// Post POST verb
func (d *DocumentHandleJSON) Post(payload string) HTTPResponse {
	status := http.StatusMethodNotAllowed
	response := "POST method not allowed json document is not a collection"

	if d.Document.IsCollection {
		var item map[string]interface{}
		err := json.Unmarshal([]byte(payload), &item)
		if err == nil {
			rawCollection := *d.Document.ResponseBody
			collection := rawCollection.([]interface{})
			collection = append(collection, item)
			rawCollection = collection
			d.Document.ResponseBody = &rawCollection

			status = http.StatusOK
			response = "success"
		} else {
			status = http.StatusInternalServerError
			response = "Error while parsing payload"
		}
	}

	return HTTPResponse{
		Status: status,
		Body:   response,
	}
}

//Put PUT verb
func (d *DocumentHandleJSON) Put(key string, payload string) HTTPResponse {
	response := "PUT method not allowed json document is not a collection"
	status := http.StatusMethodNotAllowed

	if d.Document.IsCollection {
		index, _ := d.search(key)
		if index > -1 {
			var item map[string]interface{}
			err := json.Unmarshal([]byte(payload), &item)

			if err == nil {
				rawCollection := *d.Document.ResponseBody
				collection := rawCollection.([]interface{})
				collection[index] = item
				rawCollection = collection
				d.Document.ResponseBody = &rawCollection
				status = http.StatusOK
				response = "success"
			} else {
				status = http.StatusInternalServerError
				response = "Error while parsing payload"
			}
		} else {
			status = http.StatusNotFound
			response = "Item not Found"
		}
	}

	return HTTPResponse{
		Status: status,
		Body:   response,
	}
}

func (d DocumentHandleJSON) toString(object interface{}) (string, error) {
	rawJSON, err := json.Marshal(object)
	return string(rawJSON), err
}

func (d DocumentHandleJSON) toJSON(rawJSON string) (interface{}, error) {
	var item map[string]interface{}
	err := json.Unmarshal([]byte(rawJSON), &item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (d *DocumentHandleJSON) search(key string) (int, interface{}) {
	var collection interface{}
	collection = *d.Document.ResponseBody
	if key != "" {
		for i, v := range collection.([]interface{}) {
			m := v.(map[string]interface{})
			value := fmt.Sprintf("%v", m[d.Document.FieldMatcher])
			if value == key {
				return i, v
			}
		}
	}

	return -1, nil
}

func (d *DocumentHandleJSON) remove(index int) {
	rawCollection := *d.Document.ResponseBody
	collection := rawCollection.([]interface{})
	collection = append(collection[:index], collection[index+1:]...)
	rawCollection = collection
	d.Document.ResponseBody = &rawCollection
}
