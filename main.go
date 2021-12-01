package main

import (
	"fmt"
)

//global variables
var holes []Hole
var text string
var originalText string
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

		// Print function is used to
		// display output in the same line
		//fmt.Print("You have selected: ")

		// Addition of two string
		//fmt.Println(topic)
		fmt.Println("Please enter a topic for your madlib. Ex: penguins")

		fmt.Scanln(&searchTerm)

		fmt.Println("Great! Generating a madlib from " + topic + " about " + searchTerm + "...")

		text = Scrape(topic, searchTerm)
		originalText = text
		//fmt.Println(Scrape(topic, searchTerm))

		//this is just example text until we implement a method that finds text for us
		//text = "This is how the birth of Jesus the Messiah came about: His mother Mary was pledged to be married to Joseph, but before they came together, she was found to be pregnant through the Holy Spirit. Because Joseph her husband was faithful to the law, and yet did not want to expose her to public disgrace, he had in mind to divorce her quietly. But after he had considered this, an angel of the Lord appeared to him in a dream and said, “Joseph son of David, do not be afraid to take Mary home as your wife, because what is conceived in her is from the Holy Spirit. She will give birth to a son, and you are to give him the name Jesus, because he will save his people from their sins.” All this took place to fulfill what the Lord had said through the prophet: “The virgin will conceive and give birth to a son, and they will call him Immanuel” (which means “God with us”). When Joseph woke up, he did what the angel of the Lord had commanded him and took Mary home as his wife. But he did not consummate their marriage until she gave birth to a son. And he gave him the name Jesus."
		parseText()

		var newWords []string
		for _, element := range holes {
			fmt.Println("Please enter ", element.PartOfSpeech)
			var newWord string
			fmt.Scanln(&newWord)

			newWords = append(newWords, newWord)
		}

		insertWords(newWords)

		fmt.Println()
		fmt.Println(text)

		fmt.Println()
		fmt.Println("Would you like to see the original text? Enter y or n")
		var seeOriginal string
		fmt.Scanln(&seeOriginal)
		if seeOriginal == "y" {
			fmt.Println(originalText)
		}

		fmt.Println()
		fmt.Println("Enter p to play again or enter q to quit")
		fmt.Scanln(&playAgain)
	}
}
