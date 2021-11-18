package main

import (
	"fmt"
	"strings"
)

func insertWords() {

	var pos = 0

	var index int

	for _, element := range wordsToReplace {
		index = findWordLocation(element, indicesToReplace[pos])

		var wordlen = len(wordsToReplace[pos])

		text = text[0:index] + userInputWords[pos] + text[index+wordlen:]
		pos++
	}

	fmt.Println(text)

}

func findWordLocation(substr string, position int) int {

	var textToSearch = text[position:]
	var loc = strings.Index(textToSearch, substr)

	return loc + position

}
