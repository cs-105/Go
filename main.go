package main

import (
	"fmt"
	"log"

	htgotts "github.com/hegedustibor/htgo-tts"
)

//global variables
var posToLong map[string]string
var posToShort map[string]string

// main function
func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak("You are an awesome golang programmer.")

	colorBlue := "\033[34m"

	//populate maps
	posToLong = makePosToLong()
	posToShort = makePosToShort()

	fmt.Print(colorBlue)

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
		fmt.Println("\nPlease select a source for your madlib: ")
		fmt.Println("1 - lyrics \t 2 - news \t 3 - wikipedia")

		// var then variable name then variable type
		var topic string
		var searchTerm string

		// Taking input from user
		fmt.Scanln(&topic)
		for topic != "1" && topic != "2" && topic != "3" {
			fmt.Println("\nInvalid input.")
			fmt.Println("Please select a source for your madlib: ")
			fmt.Println("1 - lyrics \t 2 - news \t 3 - wikipedia")
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

		texts := make(chan pair)
		go Scrape(texts, topic, searchTerm)
		// one := <- texts
		// two := <- texts
		// three := <- texts
		var originalText string
		for pair := range texts {
			if pair.topic == topic {
				originalText = pair.text
			}
		}

		// taliasMap := make(map[string]string, 3)
		// taliasMap[one.topic] = one.text
		// taliasMap[two.topic] = two.text
		// taliasMap[three.topic] = three.text
		//fmt.Println(Scrape(topic, searchTerm))

		var holes []Hole = parseText(originalText)

		var newWords []string
		for _, element := range holes {
			fmt.Println("Please enter", element.PartOfSpeech)
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
