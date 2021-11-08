package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ReadHTML(c *colly.Collector) {
	
	c.MaxDepth(1)

	// Find and visit all links
	/*c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	}) */

	c.OnHTML("body", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Print(text)
	})

	c.OnRequest(func(response *colly.Request) {
		fmt.Println("Visiting", response.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		myString := string(response.Body[:])
		fmt.Println(myString)
	})

	c.Visit(str)
}
