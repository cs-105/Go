package main

import (
	"fmt"

	g "github.com/serpapi/google-search-results-golang"
)

func main() {
	api := "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320"
	parameters := map[string]string{
		"engine":  "google",
		"q":       "dog",
		"api_key": "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320",
	}

	query := g.NewGoogleSearch(parameters, api)
	results, err := query.GetJSON()
	if err != nil {
		fmt.Println(err)
	}
	organicresults := results["organic_results"].([]interface{})
	first_result := organicresults[0].(map[string]interface{})
	fmt.Println(first_result["title"].(string))
}
