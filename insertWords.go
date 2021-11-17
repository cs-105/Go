package main

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/jdkato/prose.v2"
)

func insertWords() {

	var pos = 0

	for _, element := range wordsToReplace {
		var index int = findWordLocation(element, indicesToReplace[pos])
		pos++

		fmt.Println(element)
		fmt.Println(index)
		fmt.Println(text[index:])
	}

}

func findWordLocation(substr string, position int) int {

	fmt.Println("in findWordLocation")
	fmt.Println("substr: " + substr)

	doc, err := prose.NewDocument(text)
	if err != nil {
		log.Fatal(err)
	}

	var textToSearch = text

	var found = false

	var loc int

	for found == false {
		loc = strings.Index(textToSearch, substr)
		fmt.Println(loc)

		if loc == -1 {
			fmt.Println("not found")
			found = true
		} else if loc < 20 && position < 10 {
			found = true
		} else if (strings.Index(text[loc-20:loc], doc.Tokens()[position-1].Text)) != -1 {
			found = true
		} else {
			textToSearch = textToSearch[loc:]
		}

	}

	return loc

}
