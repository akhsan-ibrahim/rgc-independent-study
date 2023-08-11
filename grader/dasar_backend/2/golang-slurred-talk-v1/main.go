package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {
	str := *words
	char := "rsz"
	for i := 0; i < len(str); i++ {
		if strings.Contains(char, string(str[i])) {
			str = strings.ReplaceAll(str, string(str[i]), "l")
		} else if strings.Contains(strings.ToUpper(char), string(str[i])) {
			str = strings.ReplaceAll(str, string(str[i]), "L")
		}
	}
	*words = str
	// TODO: answer here
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println(words)
}
