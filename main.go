package main

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
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
