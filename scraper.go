package main

type pair struct {
	topic string
	text  string
	err   error
}

func Scrape(texts chan pair, topic, searchTerm string) {

	go NewsRetriever(searchTerm, texts)
	go WikiRetriever(searchTerm, texts)
	go LyricRetriever(searchTerm, texts)

}
