package main

import (
	"strings"

	"github.com/gocolly/colly"
)

var p string

func WikiRetriever(ch chan string, searchTerm string) chan string {
	search := strings.Split(searchTerm, " ")
	query := strings.Join(search, "_")

	c := colly.NewCollector(
		colly.UserAgent("MadLibs"),
		colly.AllowURLRevisit(),
		colly.MaxDepth(1),
	)

	c.OnHTML(".mw-parser-output", func(h *colly.HTMLElement) {
		p = h.ChildText("p")
	})

	c.Visit("https://en.wikipedia.org/wiki/" + query)

	ch <- p
	return ch
}
