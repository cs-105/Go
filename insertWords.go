package main

import (
	"fmt"
	"strings"
)

func insertWords(newWords []string) {

	var index int

	var counter = 0
	for _, element := range holes {
		index = findWordLocation(element.OldWord, element.Index)
		var wordlen = len(element.OldWord)
		fmt.Println(newWords[counter])
		text = text[0:index] + newWords[counter] + text[index+wordlen:]
		counter++
	}

}

func findWordLocation(substr string, position int) int {

	var textToSearch = text[position:]
	var loc = strings.Index(textToSearch, substr)

	return loc + position

}
