package main

import "sync"

type texts struct {
	text string
	err  error
}

// scrapes the web for the search term
func Scrape(searchTerm string) (texts, texts, texts) {
	lyrics := texts{}
	wikipedia := texts{}
	news := texts{}

	// scrapes all three retrievers at the same time
	var wg sync.WaitGroup
	wg.Add(3)
	go lyrics.LyricRetriever(searchTerm, &wg)
	go news.NewsRetriever(searchTerm, &wg)
	go wikipedia.WikiRetriever(searchTerm, &wg)

	wg.Wait()

	return lyrics, news, wikipedia
}
