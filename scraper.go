package retriever

import (
	"strings"
)

var api = "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320"

func Scrape(topic, searchTerm string) {
	topic = strings.ToLower(topic)

	switch topic {
	case "wiki":
		WikiRetriever(searchTerm)
	case "news":
		NewsRetriever(searchTerm)
	case "songs":
		LyricRetriever(searchTerm)
	}

}
