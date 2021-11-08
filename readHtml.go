package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ReadHTML(str string) {
	c := colly.NewCollector(
		colly.UserAgent("MadLibs"),
		colly.MaxDepth(1),
		colly.AllowURLRevisit(),
		colly.Async(true),
	)

	// Find and visit all links
	/*c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	}) */

	c.OnHTML("head", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Print(text)
	})

	c.OnHTML("section", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Print(text)
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Println(text)
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
