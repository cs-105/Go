
   
/*
 * Talia Bjelland 10/12/21
 * Purpose: demonstrate part of speech tagging using golang
 */
package main

import (
	"fmt"
	"log"

	"gopkg.in/jdkato/prose.v2"
)

func main() {
	// Create a new document with the default configuration:
	doc, err := prose.NewDocument("18 This is how the birth of Jesus the Messiah came about[d]: His mother Mary was pledged to be married to Joseph, but before they came together, she was found to be pregnant through the Holy Spirit. 19 Because Joseph her husband was faithful to the law, and yet[e] did not want to expose her to public disgrace, he had in mind to divorce her quietly. 20 But after he had considered this, an angel of the Lord appeared to him in a dream and said, “Joseph son of David, do not be afraid to take Mary home as your wife, because what is conceived in her is from the Holy Spirit. 21 She will give birth to a son, and you are to give him the name Jesus,[f] because he will save his people from their sins.” 22 All this took place to fulfill what the Lord had said through the prophet: 23 “The virgin will conceive and give birth to a son, and they will call him Immanuel”[g] (which means “God with us”). 24 When Joseph woke up, he did what the angel of the Lord had commanded him and took Mary home as his wife. 25 But he did not consummate their marriage until she gave birth to a son. And he gave him the name Jesus.")
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag)
	}
}
