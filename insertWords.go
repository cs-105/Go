package main

import (
	"fmt"
	"strings"
)

func insertWords() {

	var index int

	for _, element := range holes {
		index = findWordLocation(element.OldWord, element.Index)

		var wordlen = len(element.OldWord)

		fmt.Println(text)

		text = text[0:index] + element.NewWord + text[index+wordlen:]
	}

	fmt.Println(text)

}

func findWordLocation(substr string, position int) int {

	var textToSearch = text[position:]
	var loc = strings.Index(textToSearch, substr)

	return loc + position

}
