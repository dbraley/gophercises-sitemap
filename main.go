package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a site map for")
	flag.Parse()

	fmt.Println(*urlFlag)

	// GET the webpage
	// Parse all the links on the page
	// build proper urls with our links
	// filter out links w/ a different name
	// find all pages (BFS)
	// print out xml
}
