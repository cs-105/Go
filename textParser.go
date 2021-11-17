package main

import (
	"log"

	"math/rand"
	"time"

	"gopkg.in/jdkato/prose.v2"
)

//helper method to see if a list of strings contains a string
func containsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//helper method to see if a list of ints contains an int
func containsInt(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//helper method creates a map to convert between symbol and part of speech
func makePosToShort() map[string]string {
	posConverter := make(map[string]string)

	posConverter["adjective"] = "JJ"
	posConverter["noun"] = "NN"
	posConverter["proper noun"] = "NNP"
	posConverter["plural noun"] = "NNS"
	posConverter["adverb"] = "RB"
	posConverter["verb"] = "VB"
	posConverter["past tense verb"] = "VBD"
	posConverter["verb non 3rd person singular present"] = "VBP"
	posConverter["verb 3rd person singular present"] = "VBZ"

	return posConverter
}

//helper method creates a map to convert between symbol and part of speech
func makePosToLong() map[string]string {
	posConverter := make(map[string]string)

	posConverter["JJ"] = "adjective"
	posConverter["NN"] = "noun"
	posConverter["NNP"] = "proper noun"
	posConverter["NNS"] = "plural noun"
	posConverter["RB"] = "adverb"
	posConverter["VB"] = "verb"
	posConverter["VBD"] = "past tense verb"
	posConverter["VBP"] = "verb non 3rd person singular present"
	posConverter["VBZ"] = "verb 3rd person singular present"

	return posConverter
}

//iterates over document
//NOTE TO SELF: YOU CAN ITERATE BY SENTENCE
func parseText() {

	// Create a new document with the default configuration:
	doc, err := prose.NewDocument(text)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	var startingIndex = 0
	canBeReplaced := []string{"JJ", "NN", "NNP", "NNS", "RB", "VB", "VBD", "VBP", "VBZ"}

	// Iterate over the doc's sentences:
	for _, sent := range doc.Sentences() {

		var oneOrTwo = rand.Intn(2) + 1
		var numFound = 0
		var attempts = 0

		doc2, err := prose.NewDocument(sent.Text)
		if err != nil {
			log.Fatal(err)
		}

		for numFound != oneOrTwo && attempts < 5 {
			var rando = rand.Intn(len(doc2.Tokens()))
			var randoTok = doc2.Tokens()[rando]

			if containsString(canBeReplaced, randoTok.Tag) && !containsInt(indicesToReplace, startingIndex+rando) {
				indicesToReplace = append(indicesToReplace, startingIndex+rando)
				wordsToReplace = append(wordsToReplace, doc2.Tokens()[rando].Text)
				posToReplace = append(posToReplace, posToLong[doc2.Tokens()[rando].Tag])
				numFound = numFound + 1
				attempts = attempts + 1
			}

		}

		startingIndex = startingIndex + len(doc2.Tokens())
	}

}
