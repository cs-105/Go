package main

import (
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

var p string

// retrieves a wiki page
func (text *texts) WikiRetriever(searchTerm string, wg *sync.WaitGroup) {

	// reformats search into acceptable query
	search := strings.Split(searchTerm, " ")
	query := strings.Join(search, "_")
	para := make([]string, 2)

	// initializes new collector
	c := colly.NewCollector(
		colly.UserAgent("MadLibs"),
		colly.AllowURLRevisit(),
		colly.MaxDepth(1),
	)

	pos := 0

	// finds proper selector
	go c.OnHTML(".mw-parser-output", func(h *colly.HTMLElement) {
		// iterates through every paragraph element
		h.ForEachWithBreak("p", func(i int, h *colly.HTMLElement) bool {
			// collects paragraphs 3-5 of the page
			if i == 5 {
				return false
			} else if i > 2 {
				para[pos] = h.Text
				pos++
			}
			return true
		})
	})

	// visits the searched page
	err := c.Visit("https://en.wikipedia.org/wiki/" + query)

	if err != nil {
		text.err = err
		text.text = ""
	} else {
		p = strings.Join(para, " ")
		p = strings.Trim(p, "\n")
		text.text = p
		text.err = err
	}
	wg.Done()
}
