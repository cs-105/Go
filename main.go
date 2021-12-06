/*
 * Written by Ethan Banez, Michael Todd, Talia Bjelland, 2021
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// main function
func main() {

	//clear audio folder in preparation for text to speech
	err2 := os.RemoveAll("audio")
	if err2 != nil {
		log.Fatal(err2)
	}

	blue()

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

		//get topic and term from user
		colorReset()
		fmt.Println("\nPlease select a source for your madlib: ")
		fmt.Println("1 - lyrics \t 2 - news \t 3 - wikipedia")
		var topic string
		var searchTerm string

		blue()
		fmt.Scanln(&topic)
		for topic != "1" && topic != "2" && topic != "3" {
			red()
			fmt.Println("\nInvalid input.")
			colorReset()
			fmt.Println("Please select a source for your madlib: ")
			fmt.Println("1 - lyrics \t 2 - news \t 3 - wikipedia")
			blue()
			fmt.Scanln(&topic)
		}

		if topic == "1" {
			topic = "lyrics"
		} else if topic == "2" {
			topic = "news"
		} else if topic == "3" {
			topic = "wikipedia"
		}

		colorReset()
		fmt.Println("Please enter a topic for your madlib. Ex: penguins")

		//fmt.Scanln(&searchTerm)
		//to read in spaces as well
		blue()
		in := bufio.NewReader(os.Stdin)
		searchTerm, _ = in.ReadString('\n')
		// trims off the newline character
		searchTerm = strings.TrimSpace(searchTerm)

		colorReset()
		fmt.Println("Great! Generating a madlib from " + topic + " about " + searchTerm + "...")

		//notify user about success of search
		lyrics, news, wikipedia := Scrape(searchTerm)

		options := make(map[string]texts)

		options["lyrics"] = lyrics
		options["news"] = news
		options["wikipedia"] = wikipedia

		for option := range options {
			if options[option].err != nil {
				delete(options, option)
			}
		}

		var originalText string

		if _, ok := options[topic]; ok {
			originalText = options[topic].text
		} else {
			fmt.Printf("There were no results for %s\n", topic)
			for option := range options {
				fmt.Printf("You can try '%s' for %s\n", option, option)
			}
			fmt.Scanln(&topic)
			originalText = options[topic].text
		}

		text := originalText

		//choose madlib holes from text
		var holes []Hole = parseText(originalText)

		//get user input for each hole
		var newWords []string
		for _, element := range holes {
			colorReset()
			fmt.Println("Please enter", element.PartOfSpeech)
			var newWord string
			blue()
			fmt.Scanln(&newWord)
			newWords = append(newWords, newWord)
		}

		//put user inputted words back in text
		text = insertWords(newWords, holes, text)

		colorReset()
		fmt.Println("\n" + text)

		//play text to speech files
		playSound(text)

		colorReset()
		fmt.Println("\nWould you like to see the original text? Enter y or n")
		var seeOriginal string
		blue()
		fmt.Scanln(&seeOriginal)
		colorReset()
		if seeOriginal == "y" {
			fmt.Println(originalText)
		}

		fmt.Println("\nEnter p to play again or enter q to quit")

		blue()
		fmt.Scanln(&playAgain)
		colorReset()
	}

}
