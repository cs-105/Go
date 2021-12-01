package main

import (
	"log"
	"strings"
	"sync"

	"gopkg.in/jdkato/prose.v2"
)

func parseText() {

	// Create a new document with the default configuration:
	doc, err := prose.NewDocument(text)
	if err != nil {
		log.Fatal(err)
	}

	canBeReplaced := []string{"JJ", "NN", "NNP", "NNS", "RB", "VB", "VBD", "VBP", "VBZ"}

	var wg sync.WaitGroup
	wg.Add(len(doc.Tokens()))

	var possibleReplacers []prose.Token

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

	for i := 0; i < len(possibleReplacers); i = i + 3 {
		var hole = new(Hole)
		hole.Index = 0
		hole.OldWord = possibleReplacers[i].Text
		hole.PartOfSpeech = posToLong[possibleReplacers[i].Tag]
		holes = append(holes, *hole)
	}

}
