package main

import (
	"bufio"
	"log"
	"os"
)



func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	f1, err := os.Open("sample.txt")
	check(err)
	defer f1.Close()

	f2, err := os.Create("result.txt")
	check(err)
	defer f2.Close()

	text := bufio.NewScanner(f1)
	writer := bufio.NewWriter(f2)

	for text.Scan() {
		line := text.Text()


		new := articleA(line)
		new = Capitalize(new)
		new = Allhex(new)
		new = Quotes(new)


		_, err := writer.WriteString(new + "\n")
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()
	}
	if err := text.Err(); err != nil {
		log.Fatal(err)
	}
}
