package main

import (
	"math/rand"
	"strings"

	"github.com/gocolly/colly"
)

var coll = colly.NewCollector(
	colly.UserAgent("MadLibs"),
	colly.MaxDepth(1),
	colly.AllowURLRevisit(),
)
var newsLinks = make(map[string]int)

func NewsRetriever(searchTerm string) string {
	linkRetriever(searchTerm, newsLinks)
	link := linkFollower(newsLinks)
	text := pageReader(link)
	return text
}

func linkRetriever(searchTerm string, links map[string]int) map[string]int {
	c := colly.NewCollector(
		colly.UserAgent("MadLibs"),
		colly.MaxDepth(1),
		colly.AllowURLRevisit(),
	)

	news := c.Clone()
	linkCounter := 0

	search := strings.Split(searchTerm, " ")
	query := strings.Join(search, "+")
	// Find and visit all links

	c.OnHTML(".views-row.views-row-1.views-row-odd.views-row-first.search-value", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		news.Visit("https://www.allsides.com" + link)
	})

	news.OnHTML(".news-item.center", func(e *colly.HTMLElement) {
		linkCounter++
		link := e.ChildAttr("a", "href")
		links[link] = linkCounter
	})

	news.OnHTML(".news-item.right", func(e *colly.HTMLElement) {
		linkCounter++
		link := e.ChildAttr("a", "href")
		links[link] = linkCounter
	})

	news.OnHTML(".news-item.right", func(e *colly.HTMLElement) {
		linkCounter++
		link := e.ChildAttr("a", "href")
		links[link] = linkCounter
	})

	c.Visit("https://www.allsides.com/search?search=" + query)

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
	coll.OnHTML(".article-description", func(e *colly.HTMLElement) {
		text = e.ChildText("p")
	})

	coll.Visit(link)
	return text
}
