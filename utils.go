/*
* Written by Talia Bjelland, 2021
* Purpose: collection of small helper methods that could be reused in multiple contexts
 */
package main

import "fmt"

type Hole struct {
	PartOfSpeech string
	Index        int
	OldWord      string
	NewWord      string
}

var posToLong map[string]string = map[string]string{"JJ": "an adjective", "NN": "a noun", "NNP": "a proper noun", "NNS": "a plural noun", "RB": "an adverb", "VB": "a verb", "VBD": "a past tense verb"}

//helper method to see if a list of strings contains a string
func containsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//sets output color to blue
func blue() {
	colorBlue := "\033[34m"
	fmt.Print(string(colorBlue))
}

//resets output color to default
func colorReset() {
	colorReset := "\033[0m"
	fmt.Print(string(colorReset))
}

//sets output color to red
func red() {
	colorRed := "\033[31m"
	fmt.Print(string(colorRed))
}
