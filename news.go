package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	linkRetriever("Joe Biden")
}

func linkRetriever(searchTerm string) {
	c := colly.NewCollector(
		colly.UserAgent("MadLibs"),
		colly.MaxDepth(1),
		colly.AllowURLRevisit(),
	)

	// Find and visit all links
	/*c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	}) */

	c.OnHTML(".news-item.center", func(e *colly.HTMLElement) {
		e.ChildText(".news-item a")
	})

	c.OnHTML(".news-item.right", func(e *colly.HTMLElement) {
		e.ChildText(".news-item a")
	})

	c.OnHTML(".news-item.left", func(e *colly.HTMLElement) {
		e.ChildText(".news-item a")
	})

	c.OnRequest(func(response *colly.Request) {
		fmt.Println("Visiting", response.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		myString := string(response.Body[:])
		fmt.Println(myString)
	})

	c.Visit("https://www.allsides.com/search?search=Joe+Biden")

}
