package main

import (
	"log"

	"math/rand"
	"time"

	"gopkg.in/jdkato/prose.v2"
)

//iterates over document
//NOTE TO SELF: YOU CAN ITERATE BY SENTENCE
func parseText() {

	// Create a new document with the default configuration:
	doc, err := prose.NewDocument(text)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	var startingIndex = 0
	canBeReplaced := []string{"JJ", "NN", "NNP", "NNS", "RB", "VB", "VBD", "VBP", "VBZ"}

	// Iterate over the doc's sentences:
	for _, sent := range doc.Sentences() {

		sentDoc, err := prose.NewDocument(sent.Text)
		if err != nil {
			log.Fatal(err)
		}

		var possibleReplacers []prose.Token

		for _, tok := range sentDoc.Tokens() {
			if len(tok.Text) > 4 && containsString(canBeReplaced, tok.Tag) {
				possibleReplacers = append(possibleReplacers, tok)
			}
		}

		var twoOrThree = rand.Intn(2) + 2
		var numFound = 0
		var alreadyFound []string

		for numFound != twoOrThree && numFound < len(possibleReplacers) {
			var rando = rand.Intn(len(possibleReplacers))
			if !containsString(alreadyFound, possibleReplacers[rando].Text) {

				var hole = new(Hole)
				hole.Index = startingIndex
				hole.OldWord = possibleReplacers[rando].Text
				hole.PartOfSpeech = posToLong[possibleReplacers[rando].Tag]
				holes = append(holes, *hole)
				alreadyFound = append(alreadyFound, hole.OldWord)

				numFound++
			}
		}
		startingIndex = startingIndex + len(sent.Text)
	}

}
