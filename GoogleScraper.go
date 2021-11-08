package main

import (
	"bufio"
	"fmt"
	"os"

	g "github.com/serpapi/google-search-results-golang"
)

func main() {
	api := "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320"

	fmt.Println("What kind of articles would you like to madlib on?")
	reader := bufio.NewReader(os.Stdin)
	searchTerm, _ := reader.ReadString('\n')

	parameters := map[string]string{
		"engine":  "google",
		"q":       searchTerm,
		"output":  "html",
		"api_key": "8a54bed3f33a5d9127170bc6b3af978878ba7400e9e4c1cf3e0476fdada43320",
	}

	query := g.NewGoogleSearch(parameters, api)
	results, err := query.GetJSON()
	if err != nil {
		fmt.Println(err)
	}

	organicresults := results["organic_results"].([]interface{})
	var pageReturned map[string]interface{}
	for _, v := range organicresults {
		pageReturned = v.(map[string]interface{})
		fmt.Printf("%v: %v\n", pageReturned["title"], pageReturned["link"])
	}
}
