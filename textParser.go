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

	/*posConverter["JJ"] = "adjective"
	posConverter["NN"] = "noun"
	posConverter["NNP"] = "proper noun"
	posConverter["NNS"] = "plural noun"
	posConverter["RB"] = "adverb"
	posConverter["VB"] = "verb"
	posConverter["VBD"] = "past tense verb"
	posConverter["VBP"] = "verb non 3rd person singular present"
	posConverter["VBZ"] = "verb 3rd person singular present"*/

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

	//prepare iterating variables
	rand.Seed(time.Now().UnixNano())
	var sentence []string
	var startingIndex = 0
	var indices []int
	var index int = -1
	canBeReplaced := []string{"JJ", "NN", "NNP", "NNS", "RB", "VB", "VBD", "VBP", "VBZ"}

	for _, tok := range doc.Tokens() {
		//fmt.Println(tok.Text, tok.Tag)
		index++

		if tok.Tag == "." || tok.Text == "!" || tok.Text == "" {

			//randomly choose either one or two words to be chosen from sentence
			var oneOrTwo = rand.Intn(2) + 1
			for i := 0; i < oneOrTwo; i++ {
				var length = len(sentence)
				var rando = rand.Intn(length)
				if !containsInt(indicesToReplace, startingIndex+indices[rando]) {
					indicesToReplace = append(indicesToReplace, startingIndex+indices[rando])
					posToReplace = append(posToReplace, sentence[rando])
				}
			}
			sentence = nil
			startingIndex = index + 1
		}

		if containsString(canBeReplaced, tok.Tag) {
			sentence = append(sentence, tok.Tag)
			indices = append(indices, index)
		}
	}

}
