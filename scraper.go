package main

type texts struct {
	text string
	err  error
}

func Scrape(searchTerm string) (texts, texts, texts) {
	lyrics := texts{}
	wikipedia := texts{}
	news := texts{}

	lyrics.LyricRetriever(searchTerm)
	news.NewsRetriever(searchTerm)
	wikipedia.WikiRetriever(searchTerm)

	return lyrics, news, wikipedia
}
