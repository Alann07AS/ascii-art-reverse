package funcReverse

import (
	"log"
	"os"
	"strings"
)

func CompileAscii() map[rune]string {
	table, err := os.ReadFile("standard.txt")
	tableLn := strings.Split(string(table), "\n")
	if err != nil {
		log.Fatal(err)
	}
	charColections := make(map[rune]string)
	firstRune := ' '
	ch := ""
	for _, each := range tableLn {
		if len(each) != 0 {
			ch += each[:len(each)-1]
		} else if ch != "" {
			charColections[firstRune] = ch
			firstRune++
			ch = ""
		}
	}
	return charColections
}
