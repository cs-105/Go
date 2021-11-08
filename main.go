package main

import "fmt"

//global variables
var indicesToReplace []int
var posToReplace []string
var text string

// main function
func main() {

	fmt.Println("Welcome to Interactive Madlibs!")

	// Println function is used to
	// display output in the next line
	fmt.Println("To begn, please enter a topic: ")

	// var then variable name then variable type
	var userInput string

	// Taking input from user
	fmt.Scanln(&userInput)

	// Print function is used to
	// display output in the same line
	fmt.Print("You have selected: ")

	// Addition of two string
	fmt.Print(userInput)

	//this is just example text until we implement a method that finds text for us
	text = "This is how the birth of Jesus the Messiah came about[d]: His mother Mary was pledged to be married to Joseph, but before they came together, she was found to be pregnant through the Holy Spirit. 19 Because Joseph her husband was faithful to the law, and yet[e] did not want to expose her to public disgrace, he had in mind to divorce her quietly. 20 But after he had considered this, an angel of the Lord appeared to him in a dream and said, “Joseph son of David, do not be afraid to take Mary home as your wife, because what is conceived in her is from the Holy Spirit. 21 She will give birth to a son, and you are to give him the name Jesus,[f] because he will save his people from their sins.” 22 All this took place to fulfill what the Lord had said through the prophet: 23 “The virgin will conceive and give birth to a son, and they will call him Immanuel”[g] (which means “God with us”). 24 When Joseph woke up, he did what the angel of the Lord had commanded him and took Mary home as his wife. 25 But he did not consummate their marriage until she gave birth to a son. And he gave him the name Jesus."
	parseText()

}

//findText function should pass the topic in string form and return a relevant body of text found on the internet. This text is then parsed by textParser
//ETHAN
func findText() {

}

//inputValidation should prevent users from inputting a word of different type than the word that was removed
func inputValidation() {

}

//insertWords should take the validated user inputs and pass them back into the body of text
func insertWords() {

}
