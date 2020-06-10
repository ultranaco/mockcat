package main

import (
	"bytes"
	"fmt"
	"html/template"
	"path"
)

// Help properties to build usage message
type Help struct {
	IsCollection   bool
	CurlCollection string
	CurlItem       string
}

// New initialize help model
func (h *Help) New(port string, descriptor DocumentDescriptor) {
	h.CurlCollection = "http://localhost" + path.Clean(fmt.Sprintf("%s%s/{%s}", port, descriptor.Route, descriptor.FieldMatcher))
	h.CurlItem = fmt.Sprintf("http://localhost%s%s", port, descriptor.Route)
	h.IsCollection = descriptor.IsCollection
}

// Print usage
func (h *Help) Print() {

	t := template.Must(template.New("todos").Parse(`
Usage:

    curl {{ .CurlItem}}
		{{if .IsCollection}}
    or

    curl {{ .CurlCollection}}
		{{end}}
Commands:
    -m   :property matcher used to retrieve an item over a collection through of root properties
    -p   :port to listen incoming requests
	`))

	var buf bytes.Buffer
	t.Execute(&buf, *h)

	fmt.Printf("%s", buf.String())
}
