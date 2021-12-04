package main

import (
	"strings"
)

type pair struct {
	topic string
	text  string
}

func Scrape(texts chan pair, topic, searchTerm string) {
	topic = strings.ToLower(topic)

	go WikiRetriever(searchTerm, texts)
	go NewsRetriever(searchTerm, texts)
	go LyricRetriever(searchTerm, texts)

}
