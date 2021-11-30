package main

import (
	"strings"

	"github.com/gocolly/colly"
)

var book string

func bookretriever(searchTerm string) string {

	split := strings.Split(searchTerm, " ")
	query := strings.Join(split, "+")

	var c = colly.NewCollector(
		colly.MaxDepth(1),
		colly.UserAgent("MadLibs"),
	)

	var d = c.Clone()
	var e = c.Clone()

	c.OnHTML(".grayed.navlink", func(h *colly.HTMLElement) {
		book = h.ChildText("span")
		return
	})

	c.OnHTML(".booklink", func(h *colly.HTMLElement) {
		d.Visit(h.ChildAttr("href", "a"))
	})

	d.OnHTML("[title|=Plain Text", func(h *colly.HTMLElement) {
		e.Visit(h.Attr("href"))
	})

	e.OnHTML("body", func(h *colly.HTMLElement) {
		book = h.ChildText("pre")
	})

	c.Visit("https://www.gutenberg.org/ebooks/search/?query=" + query)

	return book
}
