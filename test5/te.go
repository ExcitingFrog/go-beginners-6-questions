package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// link := e.Attr("href")
	// c.Visit(e.Request.AbsoluteURL(link))
	// })
	//
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		println(string(r.Body))
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://tumregels.github.io/Network-Programming-with-Go/")
}
