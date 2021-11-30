package main

import (
	"fmt"
)

//global variables
var holes []Hole
var text string
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

	fmt.Println("Welcome to Interactive Madlibs!")

	// Println function is used to
	// display output in the next line
	fmt.Println("To begin, please enter a topic: ")

	// var then variable name then variable type
	var userInput string

	// Taking input from user
	fmt.Scanln(&userInput)

	// Print function is used to
	// display output in the same line
	fmt.Print("You have selected: ")

	// Addition of two string
	fmt.Println(userInput)

	fmt.Println(Scrape("news", userInput))

	//this is just example text until we implement a method that finds text for us
	text = "This is how the birth of Jesus the Messiah came about: His mother Mary was pledged to be married to Joseph, but before they came together, she was found to be pregnant through the Holy Spirit. Because Joseph her husband was faithful to the law, and yet did not want to expose her to public disgrace, he had in mind to divorce her quietly. But after he had considered this, an angel of the Lord appeared to him in a dream and said, “Joseph son of David, do not be afraid to take Mary home as your wife, because what is conceived in her is from the Holy Spirit. She will give birth to a son, and you are to give him the name Jesus, because he will save his people from their sins.” All this took place to fulfill what the Lord had said through the prophet: “The virgin will conceive and give birth to a son, and they will call him Immanuel” (which means “God with us”). When Joseph woke up, he did what the angel of the Lord had commanded him and took Mary home as his wife. But he did not consummate their marriage until she gave birth to a son. And he gave him the name Jesus."
	parseText()

	var newWords []string
	for _, element := range holes {
		fmt.Println("Please enter ", element.PartOfSpeech)
		var newWord string
		fmt.Scanln(&newWord)
		fmt.Println("You have entered: ", newWord)

		newWords = append(newWords, newWord)
	}

	insertWords(newWords)

	fmt.Println(text)
}

//findText function should pass the topic in string form and return a relevant body of text found on the internet. This text is then parsed by textParser
//ETHAN
func findText() {

}
