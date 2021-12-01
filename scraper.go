package main

import (
	"strings"
)

var api = "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320"

func Scrape(topic, searchTerm string) string {
	topic = strings.ToLower(topic)
	var result string
	switch topic {
	case "wiki":
		result = WikiRetriever(searchTerm)
	case "news":
		result = NewsRetriever(searchTerm)
	case "lyrics":
		result = LyricRetriever(searchTerm)
	}
	return result
}
