package main

import (
	"log"

	"gopkg.in/jdkato/prose.v2"
)

//currently not in action because it's too picky
func validate(expected string, actual string) bool {
	//use prose to get actual pos
	doc, err := prose.NewDocument(actual)
	if err != nil {
		log.Fatal(err)
	}

	var expectedAsShort = posToShort[expected] //okay map this
	return (doc.Tokens()[0].Tag == expectedAsShort)
}
