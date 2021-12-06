package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	g "github.com/serpapi/google-search-results-golang"
)

// retrieves lyrics using google api
func (texts *texts) LyricRetriever(searchTerm string, wg *sync.WaitGroup) {
	api := "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320"
	parameters := map[string]string{
		"engine":  "google",
		"q":       searchTerm + " lyrics",
		"output":  "json",
		"api_key": "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320",
	}

	query := g.NewGoogleSearch(parameters, api)
	results, err := query.GetJSON()
	if err != nil {
		fmt.Println(err)
	}

	// if there is a lyric box then we grab it
	// else return an error
	if answerBox, ok := results["answer_box"].(map[string]interface{}); ok {
		// splits on the newline characters
		lines := strings.Split(answerBox["lyrics"].(string), "\n")

		// collects only the first 8 lines
		// rejoins it in it's original format
		lines = lines[:16]
		lyrics := strings.Join(lines, "\n")

		texts.text = lyrics
		texts.err = nil
	} else {
		err := errors.New("No lyrics found.")
		texts.text = ""
		texts.err = err
	}
	wg.Done()
}
