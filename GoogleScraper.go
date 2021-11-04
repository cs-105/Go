package main

import (
	g "github.com/serpapi/google-search-results-golang"
)

func results() {
	api := "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320"
	parameters := map[string]string{
		"engine":  "google",
		"q":       "dog",
		"api_key": "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320",
	}

	query := g.NewGoogleSearch(parameters, api)
	results, err := search.GetJSON()
	organic_results := results["organic_results"].([]interface{})
}
