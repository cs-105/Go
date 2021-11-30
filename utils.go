package main

type Hole struct {
	PartOfSpeech string
	Index        int
	OldWord      string
	NewWord      string
}

//helper method to see if a list of strings contains a string
func containsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//helper method creates a map to convert between symbol and part of speech
func makePosToShort() map[string]string {
	posConverter := make(map[string]string)

	posConverter["an adjective"] = "JJ"
	posConverter["a noun"] = "NN"
	posConverter["a proper noun"] = "NNP"
	posConverter["a plural noun"] = "NNS"
	posConverter["an adverb"] = "RB"
	posConverter["a verb"] = "VB"
	posConverter["a past tense verb"] = "VBD"
	posConverter["a verb non 3rd person singular present"] = "VBP"
	posConverter["a verb 3rd person singular present"] = "VBZ"

	return posConverter
}

//helper method creates a map to convert between symbol and part of speech
func makePosToLong() map[string]string {
	posConverter := make(map[string]string)

	posConverter["JJ"] = "an adjective"
	posConverter["NN"] = "a noun"
	posConverter["NNP"] = "a proper noun"
	posConverter["NNS"] = "a plural noun"
	posConverter["RB"] = "an adverb"
	posConverter["VB"] = "a verb"
	posConverter["VBD"] = "a past tense verb"
	posConverter["VBP"] = "a verb non 3rd person singular present"
	posConverter["VBZ"] = "a verb 3rd person singular present"

	return posConverter
}
