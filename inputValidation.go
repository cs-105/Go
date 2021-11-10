package main

import (
	"log"

	"gopkg.in/jdkato/prose.v2"
)

func validate(expected string, actual string) bool {
	//use prose to get actual pos
	doc, err := prose.NewDocument(actual)
	if err != nil {
		log.Fatal(err)
	}

	var expectedAsShort = posToShort[expected] //okay map this
	// Iterate over the doc's tokens:
	return (doc.Tokens()[0].Tag == expectedAsShort)

	//use map to check expected pos
	//return true or false
}