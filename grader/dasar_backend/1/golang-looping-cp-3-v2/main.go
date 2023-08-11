package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	var result int

	for _, letter := range strings.ToUpper(text) {
		if string(letter) == "R" || string(letter) == "S" || string(letter) == "T" || string(letter) == "Z" {
			result += 1
		}
	}
	return result // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
