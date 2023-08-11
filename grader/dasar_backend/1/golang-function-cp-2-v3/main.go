package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowel := "aiueo"
	var letter string
	var cons int
	var vow int

	for i := 65; i <= 90; i++ {
		letter += string(byte(i))
	}

	for i := 0; i < len(str); i++ {
		if strings.Contains(vowel, strings.ToLower(string(str[i]))) == false && strings.Contains(strings.ToLower(letter), strings.ToLower(string(str[i]))) == true {
			cons++
		} else if strings.Contains(vowel, strings.ToLower(string(str[i]))) == true {
			vow++
		}
	}

	var check bool
	if cons == 0 || vow == 0 {
		check = true
	}

	return vow, cons, check // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
	fmt.Println(CountVowelConsonant("bbbbb ccccc"))
	fmt.Println(CountVowelConsonant("kopi"))
}
