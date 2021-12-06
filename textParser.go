/*
* Written by Talia Bjelland, 2021
* Purpose: find suitable words to replace in madlib text
 */

package main

import (
	"log"
	"strings"
	"sync"

	"github.com/jdkato/prose/v2"
)

func parseText(text string) []Hole {

	doc, err := prose.NewDocument(text)
	if err != nil {
		log.Fatal(err)
	}

	canBeReplaced := []string{"JJ", "NN", "NNP", "NNS", "RB", "VB", "VBD"}

	var wg sync.WaitGroup
	wg.Add(len(doc.Tokens()))

	var possibleReplacers []prose.Token

	//iterate over all tokens and identify those that are suitable for replacement
	for i := 0; i < len(doc.Tokens()); i++ {
		go func(i int) {
			defer wg.Done()
			var tok = doc.Tokens()[i]
			if len(tok.Text) > 4 && containsString(canBeReplaced, tok.Tag) && strings.Index(tok.Text, ".") < 0 {
				possibleReplacers = append(possibleReplacers, doc.Tokens()[i])
			}
		}(i)
	}
	wg.Wait()

	//choose every 3rd replaceable word to be replaced
	//due to the multithreading nature of code above, they will be relatively randomly replaced throughout the file
	var holes []Hole
	for i := 0; i < len(possibleReplacers); i = i + 3 {
		var hole = new(Hole)
		hole.OldWord = possibleReplacers[i].Text
		hole.PartOfSpeech = posToLong[possibleReplacers[i].Tag]
		holes = append(holes, *hole)
	}

	return holes

}
