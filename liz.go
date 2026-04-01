package main

import (
	"strings"
)

func Quotes(s string) string {
	words := strings.Split(s, "'")
	result := "  "
	for i := 0; i < len(words); i++ {
		if i%2 == 1 {
			result += "'" + strings.TrimSpace(words[i]) + "'"
		} else {
			result += words[i]
		}
	}
	return result
}
