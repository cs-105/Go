package main

import (
	"errors"
	"fmt"

	g "github.com/serpapi/google-search-results-golang"
)

var lyrics string

func (texts *texts) LyricRetriever(searchTerm string) {
	api := "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320"
	parameters := map[string]string{
		"engine":  "google",
		"q":       searchTerm + " song lyrics",
		"output":  "json",
		"api_key": "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320",
	}

	query := g.NewGoogleSearch(parameters, api)
	results, err := query.GetJSON()
	if err != nil {
		fmt.Println(err)
	}
	if answerBox, ok := results["answer_box"].(map[string]interface{}); ok {
		texts.text = answerBox["lyrics"].(string)
		texts.err = nil
	} else {
		err := errors.New("No lyrics found.")
		texts.text = ""
		texts.err = err
	}
}
