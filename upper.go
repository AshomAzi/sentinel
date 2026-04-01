package main

import (
	"strings"
)

func Capitalize(s string) string {
	e := "(cap)"
	d := strings.Fields(s)

	for i := 0; i < len(d); i++ {
		if d[i] == e && i > 0 {
			d[i-1] = strings.Title(d[i-1])
			d = append(d[:i], d[i+1:]...)
		}
	}
	return strings.Join(d, " ")
}
