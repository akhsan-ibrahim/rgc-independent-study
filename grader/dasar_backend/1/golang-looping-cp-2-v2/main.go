package main

import (
	"fmt"
	// "strings"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var result, reverse string

	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	for i := 0; i < len(reverse); i++ {
		if i == 0 || string(reverse[i]) == " " || string(reverse[i-1]) == " " {
			result += string(reverse[i])
		} else {
			result += "_" + string(reverse[i])
		}
	}

	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
