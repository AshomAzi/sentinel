package main

import "strings"

// operation gopher protocol: Article feature Handled by Fellow Enock Victor
func articleA(text string) string {
	// text = strings.ToLower(text)
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		nextWord := words[i+1][:1]
		vowels := "aeiouh"
		consonants := "bcdfgjklmnpqrstvwxyz"
		if words[i] == "a" && strings.ContainsAny(strings.ToLower(nextWord), vowels) {
			words[i] = "an"
		}

		if words[i] == "an" && strings.ContainsAny(strings.ToLower(nextWord), consonants) {
			words[i] = "a"
		}

	}
	res := strings.Join(words, " ")
	return strings.ToUpper(string(res[:0])) + strings.ToLower(res[0:])
}
