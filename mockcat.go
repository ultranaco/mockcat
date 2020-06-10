package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	fileArg := os.Args[1]
	flagSet := flag.NewFlagSet("", flag.ExitOnError)
	fieldMatcher := flagSet.String("m", "id", "property matcher")
	port := flagSet.Int("p", 8080, "port")
	flagSet.Parse(os.Args[2:])

	p := fmt.Sprintf(":%v", *port)

	file := FileDescriptor{}
	doc := DocumentDescriptor{}
	handle := ResponseHandle{}

	file.New(fileArg)
	doc.New(file)
	doc.SetFieldMatcher(*fieldMatcher)
	handle.New(doc)

	http.HandleFunc(doc.Route, handle.Handle)

	s := &http.Server{
		Addr: p,
	}

	h := Help{}
	h.New(s.Addr, doc)
	h.Print()

	log.Fatal(s.ListenAndServe())
}
