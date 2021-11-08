package main

//Main function should output a welcome message to the user and then get user input for a topic.
//Main should call the findText function to search online for a body of text relating to the specified topic.
//main should then, for every word removed, request and retrive a word from the user of equivalent type, which are all checked by inputValue
//main then prints out the now edited body of text and asks the user if they would like to do another or quit

import (
	"fmt"
	"log"

	"math/rand"
	"time"

	"gopkg.in/jdkato/prose.v2"
)

var indicesToReplace []int
var posToReplace []string
var textAsSlice []string

//global variable text  "gopkg.in/jdkato/prose.v2"

func main() {
	/*
	  output welcome message to user //MICHAEL
	  get user input for a topic //MICHAEL
	  pass user input to findText function
	  call textParser on text to get words to replace
	  for each word to replace, output part of speech to user and call inputvalidation
	  when all words have been replace, replace words from og text with user inputted words
	  print out body of text
	  ask user if they want to repeat or quit
	*/
	textParser()
}

//findText function should pass the topic in string form and return a relevant body of text found on the internet. This text is then parsed by textParser
//ETHAN
func findText() {

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func containsInt(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//textParser should return the location and type of words that are being replaced
//TALIA
func textParser() {
	// Create a new document with the default configuration:
	doc, err := prose.NewDocument("18 This is how the birth of Jesus the Messiah came about[d]: His mother Mary was pledged to be married to Joseph, but before they came together, she was found to be pregnant through the Holy Spirit. 19 Because Joseph her husband was faithful to the law, and yet[e] did not want to expose her to public disgrace, he had in mind to divorce her quietly. 20 But after he had considered this, an angel of the Lord appeared to him in a dream and said, “Joseph son of David, do not be afraid to take Mary home as your wife, because what is conceived in her is from the Holy Spirit. 21 She will give birth to a son, and you are to give him the name Jesus,[f] because he will save his people from their sins.” 22 All this took place to fulfill what the Lord had said through the prophet: 23 “The virgin will conceive and give birth to a son, and they will call him Immanuel”[g] (which means “God with us”). 24 When Joseph woke up, he did what the angel of the Lord had commanded him and took Mary home as his wife. 25 But he did not consummate their marriage until she gave birth to a son. And he gave him the name Jesus.")
	if err != nil {
		log.Fatal(err)
	}

	canBeReplaced := []string{"JJ", "NN", "NNP", "NNS", "RB", "VB", "VBD", "VBP", "VBZ"}
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

	// Iterate over the doc's tokens:
	// for each sentence, collect all replaceable tokens
	// choose randomly between 1 and 2
	//
	rand.Seed(time.Now().UnixNano())

	var sentence []string
	var startingIndex = 0
	var indices []int
	var index int = -1

	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag)
		index++
		textAsSlice = append(textAsSlice, tok.Text)
		fmt.Println(index)

		//determine which are good

		if tok.Tag == "." {

			var oneOrTwo = rand.Intn(2) + 1

			for i := 0; i < oneOrTwo; i++ {
				var length = len(sentence)
				var rando = rand.Intn(length)
				if !containsInt(indicesToReplace, startingIndex+indices[rando]) {
					indicesToReplace = append(indicesToReplace, startingIndex+indices[rando])
					posToReplace = append(posToReplace, sentence[rando])
					fmt.Println(indicesToReplace)
					fmt.Println(posToReplace)
				}
			}
			sentence = nil
			startingIndex = index + 1
		}

		if contains(canBeReplaced, tok.Tag) {
			sentence = append(sentence, tok.Tag)
			indices = append(indices, index)
		}
	}

}

//inputValidation should prevent users from inputting a word of different type than the word that was removed
func inputValidation() {

}

//insertWords should take the validated user inputs and pass them back into the body of text
func insertWords() {

}
