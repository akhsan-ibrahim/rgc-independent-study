package main

import (
	"fmt"
	"strconv"
	// "strings"
	// "strconv"
)

func PhoneNumberChecker(number string, result *string) {
	var num string
	if number[:2] == "08" && len(number) >= 10 {
		num = number[2:]
	} else if number[:3] == "628" && len(number) >= 11 {
		num = number[3:]
	} else {
		num = "-1"
	}

	provider, _ := strconv.Atoi(string(num[:2]))
	if provider >= 11 && provider <= 15 {
		*result = "Telkomsel"
	} else if provider >= 16 && provider <= 19 {
		*result = "Indosat"
	} else if provider >= 21 && provider <= 23 {
		*result = "XL"
	} else if provider >= 27 && provider <= 29 {
		*result = "Tri"
	} else if provider >= 52 && provider <= 53 {
		*result = "AS"
	} else if provider >= 81 && provider <= 88 {
		*result = "Smartfren"
	} else {
		*result = "invalid"
	}
	// TODO: answer here
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "1234567"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
