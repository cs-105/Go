/*
* Written by Talia Bjelland, 2021
* Purpose: replace words in original madlib text with user inputted words
 */

package main

import (
	"strings"
)

func insertWords(newWords []string, holes []Hole, text string) string {

	for i, element := range holes {
		var index = strings.Index(text, element.OldWord)
		var wordlen = len(element.OldWord)
		text = text[0:index] + newWords[i] + text[index+wordlen:]
		i++
	}

	return text

}
