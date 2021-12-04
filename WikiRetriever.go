package main

import (
	"strings"

	"github.com/gocolly/colly"
)

var p string

func WikiRetriever(searchTerm string, wiki chan pair) {
	search := strings.Split(searchTerm, " ")
	query := strings.Join(search, "_")

	c := colly.NewCollector(
		colly.UserAgent("MadLibs"),
		colly.AllowURLRevisit(),
		colly.MaxDepth(1),
	)

	go c.OnHTML(".mw-parser-output", func(h *colly.HTMLElement) {
		p = h.ChildText("p")
	})

	go c.Visit("https://en.wikipedia.org/wiki/" + query)

	wiki <- pair{"wikipedia", p}
}
