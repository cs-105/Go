package main

import (
	"log"
	"strings"
	"sync"

	"github.com/jdkato/prose/v2"
)

func parseText(text string) []Hole {

	// Create a new document with the default configuration:

	doc, err := prose.NewDocument(text)
	if err != nil {
		log.Fatal(err)
	}

	canBeReplaced := []string{"JJ", "NN", "NNP", "NNS", "RB", "VB", "VBD", "VBP", "VBZ"}

	//start := time.Now()
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

	/*fmt.Print("average per token: ")
	var denom float64 = float64(int64(time.Since(start)))
	var avg float64 = (float64(len(doc.Tokens())) / denom) * 100000
	fmt.Println(avg)*/

	var holes []Hole
	for i := 0; i < len(possibleReplacers); i = i + 3 {
		var hole = new(Hole)
		hole.OldWord = possibleReplacers[i].Text
		hole.PartOfSpeech = posToLong[possibleReplacers[i].Tag]
		holes = append(holes, *hole)
	}

	return holes

}
