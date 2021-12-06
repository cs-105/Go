package main

import (
	"errors"
	"math/rand"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

// initialize collector
var coll = colly.NewCollector(
	colly.UserAgent("MadLibs"),
	colly.MaxDepth(1),
	colly.AllowURLRevisit(),
)
var newsLinks = make(map[string]int)

func (text *texts) NewsRetriever(searchTerm string, wg *sync.WaitGroup) {

	newsLinks, err := linkRetriever(searchTerm, newsLinks)

	// a bunch of error catches making sure there are no errors in any of the steps
	// but we still return if there are

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
	wg.Done()
}

// retrieves the links and adds them to a map
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

	// if there is match to the selector
	// grab the link and increment the link counter
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

// returns a random link from the map so it's not always the first one
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

// follows the final link and grabs the description
func pageReader(link string) (string, error) {
	var text string
	coll.OnHTML(".article-description", func(e *colly.HTMLElement) {
		text = e.ChildText("p")
	})

	err := coll.Visit(link)
	return text, err
}
