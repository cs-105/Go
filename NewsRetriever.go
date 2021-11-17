package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/gocolly/colly"
)

var coll *colly.Collector

func main() {
	NewsRetriever("Joe Biden")
}

func NewsRetriever(searchTerm string) string {
	var links map[string]int
	linkRetriever(searchTerm, links)
	link := linkFollower(links)
	text := pageReader(link)
	return text
}

func linkRetriever(searchTerm string, links map[string]int) map[string]int {
	c := colly.NewCollector(
		colly.UserAgent("MadLibs"),
		colly.MaxDepth(1),
		colly.AllowURLRevisit(),
		colly.Async(true),
	)

	coll = c.Clone()

	terms := strings.Split(searchTerm, " ")
	stringTerms := strings.Join(terms, "+")

	// Find and visit all links
	/*c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	}) */
	linkcounter := 0

	c.OnHTML(".news-item.center", func(e *colly.HTMLElement) {
		linkcounter++
		link := e.ChildText(".news-item a")
		links[link] = linkcounter
	})

	c.OnHTML(".news-item.right", func(e *colly.HTMLElement) {
		fmt.Println(e.DOM.Children().Text())
		linkcounter++
		link := e.ChildText(".news-item a")
		links[link] = linkcounter
	})

	c.OnHTML(".news-item.left", func(e *colly.HTMLElement) {
		linkcounter++
		link := e.ChildText(".news-item a")
		links[link] = linkcounter
	})

	/*
		c.OnRequest(func(response *colly.Request) {
			fmt.Println("Visiting", response.URL)
		})

		c.OnResponse(func(response *colly.Response) {
			myString := string(response.Body[:])
			fmt.Println(myString)
		})
	*/

	c.Visit("https://www.allsides.com/search?search=" + stringTerms)

	return links
}

func linkFollower(links map[string]int) string {
	keys := make([]string, len(links))
	i := 0
	for k := range links {
		keys[i] = k
		i++
	}

	link := keys[rand.Intn(len(keys))]

	return link
}

func pageReader(link string) string {
	var text string
	coll.OnHTML("div.article-description", func(e *colly.HTMLElement) {
		text = e.ChildText("p")
	})

	coll.Visit(link)
	return text
}
