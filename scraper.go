package main

import (
	"strings"
)

var api = "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320"

func Scrape(c chan string, topic, searchTerm string) chan string {
	topic = strings.ToLower(topic)

	switch topic {
	case "wiki":
		go WikiRetriever(c, searchTerm)
	case "news":
		go NewsRetriever(c, searchTerm)
	case "lyrics":
		go LyricRetriever(c, searchTerm)
	}

	return c
}
