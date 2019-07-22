package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	link "github.com/dbraley/gophercises-link"
)

func main() {
	urlFlag := flag.String("url", "http://gophercises.com", "the url that you want to build a site map for")
	flag.Parse()

	// GET the webpage
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()
	fmt.Println(base)

	// Parse all the links on the page
	var hrefs []string
	links, _ := link.Parse(resp.Body)

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		default:
			fmt.Println("Ignoring " + l.Href)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}

	// build proper urls with our links
	// filter out links w/ a different name
	// find all pages (BFS)
	// print out xml
}
