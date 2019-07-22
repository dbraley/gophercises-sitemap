package main

import (
	"flag"
	"io"
	"net/http"
	"os"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a site map for")
	flag.Parse()

	// GET the webpage
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	// Parse all the links on the page
	// build proper urls with our links
	// filter out links w/ a different name
	// find all pages (BFS)
	// print out xml
}
