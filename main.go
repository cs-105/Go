package main

import (
	"fmt"
	"log"
	"os"
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
		// Println function is used to
		// display output in the next line
		colorReset()
		fmt.Println("\nPlease select a source for your madlib: ")
		fmt.Println("1 - lyrics \t 2 - news \t 3 - wikipedia")

		// var then variable name then variable type
		var topic string
		var searchTerm string

		// Taking input from user
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

		blue()
		fmt.Scanln(&searchTerm)

		colorReset()
		fmt.Println("Great! Generating a madlib from " + topic + " about " + searchTerm + "...")

		lyrics, news, wikipedia := Scrape(searchTerm)

		var originalText string

		if topic == "lyrics" {
			originalText = lyrics.text
		} else if topic == "news" {
			originalText = news.text
		} else if topic == "wikipedia" {
			originalText = wikipedia.text
		}

		text := originalText

		var holes []Hole = parseText(originalText)

		var newWords []string
		for _, element := range holes {
			colorReset()
			fmt.Println("Please enter", element.PartOfSpeech)
			var newWord string
			blue()
			fmt.Scanln(&newWord)
			newWords = append(newWords, newWord)
		}

		text = insertWords(newWords, holes, text)

		colorReset()
		fmt.Println("\n" + text)

		//line 104

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

// func run(filePath string) error {
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	d, err := mp3.NewDecoder(f)
// 	if err != nil {
// 		return err
// 	}

// 	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
// 	if err != nil {
// 		return err
// 	}
// 	defer c.Close()

// 	p := c.NewPlayer()
// 	defer p.Close()

// 	fmt.Printf("Length: %d[bytes]\n", d.Length())

// 	if _, err := io.Copy(p, d); err != nil {
// 		return err
// 	}
// 	return nil
// }
