package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dmitris/deptesthuge/deptestsmall"
	"golang.org/x/net/html"
)

func main() {
	s := deptestsmall.SillyType{}.String()
	fmt.Printf("[INFO]: %s\n", s)

	res, err := http.Get("http://en.wikipedia.org")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the HTML
	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Doc: #%v\n", doc)
}
