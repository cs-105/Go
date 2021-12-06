package main

import (
	"errors"
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

func (text *texts) NewsRetriever(searchTerm string) {
	newsLinks, err := linkRetriever(searchTerm, newsLinks)
	if err != nil {
		text.text = ""
		text.err = err
	} else {
		link, err := linkFollower(newsLinks)

		if err != nil {
			text.text = link
			text.err = err
		} else {
			str, err := pageReader(link)
			text.text = str
			text.err = err
		}
	}

}

func linkRetriever(searchTerm string, links map[string]int) (map[string]int, error) {
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

	go c.OnHTML(".views-row.views-row-1.views-row-odd.views-row-first.search-value", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		news.Visit("https://www.allsides.com" + link)
	})

	go news.OnHTML(".news-item.center", func(e *colly.HTMLElement) {
		linkCounter++
		link := e.ChildAttr("a", "href")
		links[link] = linkCounter
	})

	go news.OnHTML(".news-item.right", func(e *colly.HTMLElement) {
		linkCounter++
		link := e.ChildAttr("a", "href")
		links[link] = linkCounter
	})

	go news.OnHTML(".news-item.right", func(e *colly.HTMLElement) {
		linkCounter++
		link := e.ChildAttr("a", "href")
		links[link] = linkCounter
	})

	err := c.Visit("https://www.allsides.com/search?search=" + query)

	return links, err
}

func linkFollower(links map[string]int) (string, error) {
	if len(links) != 0 {
		keys := make([]string, len(links))
		i := 0
		for k := range links {
			keys[i] = k
			i++
		}
		link := keys[rand.Intn(len(keys))]

		return link, nil
	} else {
		err := errors.New("No news found")
		return "", err
	}
}

func pageReader(link string) (string, error) {
	var text string
	coll.OnHTML(".article-description", func(e *colly.HTMLElement) {
		text = e.ChildText("p")
	})

	err := coll.Visit(link)
	return text, err
}
