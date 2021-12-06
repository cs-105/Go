package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	htgotts "github.com/hegedustibor/htgo-tts"
)

//global variables
var posToLong map[string]string
var posToShort map[string]string

// main function
func main() {

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

		//fmt.Scanln(&searchTerm)
		//to read in spaces as well
		in := bufio.NewReader(os.Stdin)
		searchTerm, _ = in.ReadString('\n')
		// trims off the newline character
		searchTerm = strings.TrimSpace(searchTerm)

		fmt.Println("Great! Generating a madlib from " + topic + " about " + searchTerm + "...")

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

		// if options[topic].err != nil {
		// 	delete(options, topic)
		// 	fmt.Printf("There were no results for %s topic.\n", topic)
		// 	for option := range options {
		// 		fmt.Printf("You can try %s\n", option)
		// 	}
		// } else {
		// 	originalText = options[topic].text
		// }

		// var options [3]texts

		// pos := 0
		// if lyrics.err != nil {
		// 	options[pos] = lyrics
		// 	pos++
		// }
		// if news.err != nil {
		// 	options[pos] = news
		// 	pos++
		// }
		// if wikipedia.err != nil {
		// 	options[pos] = wikipedia
		// }

		text := originalText

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

		//line 104
		var splitText []string
		i := 0
		j := 100
		for i < len(text) {

			if j > len(text)-1 {
				splitText = append(splitText, text[i:])
			} else {
				j = j + (strings.Index(text[j:], " "))
				splitText = append(splitText, text[i:j])
			}
			i = j
			j = j + 100
		}

		fmt.Println(text)

		speech := htgotts.Speech{Folder: "audio", Language: "en"}

		for i, element := range splitText {
			speech.Speak(element)
			files, err := ioutil.ReadDir("audio")
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range files {
				var name string = file.Name()
				if string(name[0]) == "e" {
					var dst string = "a" + strconv.Itoa(i) + ".mp3"
					os.Rename("audio/"+file.Name(), "audio/"+dst)
				}
			}

		}
		files, err := ioutil.ReadDir("audio")
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if err := run("audio/" + file.Name()); err != nil {
				log.Fatal(err)
			}
		}

		//remove all frm audio
		err2 := os.RemoveAll("audio")
		if err2 != nil {
			log.Fatal(err2)
		}

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
