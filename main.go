package main

import (
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

		fmt.Scanln(&searchTerm)

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
