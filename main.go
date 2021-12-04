package main

import (
	"fmt"
)

//global variables
var posToLong map[string]string
var posToShort map[string]string

// main function
func main() {

	//populate maps
	posToLong = makePosToLong()
	posToShort = makePosToShort()

	asciiArt :=
		`
			  /\
			 /**\
			/****\   /\
		   /      \ /**\
		  /  /\    /    \        /\    /\  /\      /\            /\/\/\  /\
		 /  /  \  /      \      /  \/\/  \/  \  /\/  \/\  /\  /\/ / /  \/  \
		/  /    \/ /\     \    /    \ \  /    \/ /   /  \/  \/  \  /    \   \
	   /  /      \/  \/\   \  /      \    /   /    \
	__/__/_______/___/__\___\__________________________________________________
	`
	fmt.Println(asciiArt)

	fmt.Println("Welcome to GO Tell it on the Mountain, an interactive madlib game!")

	var playAgain = "p"

	for playAgain == "p" {
		// Println function is used to
		// display output in the next line
		fmt.Println("Please select a source for your madlib: ")
		fmt.Println("Enter 1 for lyrics")
		fmt.Println("Enter 2 for news")
		fmt.Println("Enter 3 for wikipedia")

		// var then variable name then variable type
		var topic string
		var searchTerm string

		// Taking input from user
		fmt.Scanln(&topic)
		for topic != "1" && topic != "2" && topic != "3" {
			fmt.Println("Invalid input. Please enter 1, 2, or 3")
			fmt.Scanln(&topic)
		}

		if topic == "1" {
			topic = "lyrics"
		} else if topic == "2" {
			topic = "news"
		} else if topic == "3" {
			topic = "wikipedia"
		}

		fmt.Println("Please enter a topic for your madlib. Ex: penguins")

		fmt.Scanln(&searchTerm)

		fmt.Println("Great! Generating a madlib from " + topic + " about " + searchTerm + "...")

		var text = Scrape(topic, searchTerm)
		var originalText = text

		var holes []Hole = parseText(text)

		var newWords []string
		for _, element := range holes {
			fmt.Println("Please enter ", element.PartOfSpeech)
			var newWord string
			fmt.Scanln(&newWord)

			newWords = append(newWords, newWord)
		}

		text = insertWords(newWords, holes, text)

		fmt.Println("\n" + text)

		fmt.Println("\nWould you like to see the original text? Enter y or n")
		var seeOriginal string
		fmt.Scanln(&seeOriginal)
		if seeOriginal == "y" {
			fmt.Println(originalText)
		}

		fmt.Println("\nEnter p to play again or enter q to quit")
		fmt.Scanln(&playAgain)
	}
}
