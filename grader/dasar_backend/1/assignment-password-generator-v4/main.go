package main

import (
	"fmt"
	// "strconv"
	"strings"
	"unicode"
)

func Reverse(str string) string {
	var rev string
	for i := len(str) - 1; i >= 0; i-- {
		rev += string(str[i])
	}
	return rev // TODO: replace this
}

func Generate(str string) string {
	str = Reverse(str)
	strConv := ""

	vowelLow := "aeiou"
	vowelUp := "AEIOU"
	for i := 0; i < len(str); i++ {
		var vowIndex int
		if strings.Contains(vowelLow, string(str[i])) == true {
			for k := 0; k < len(vowelLow); k++ {
				if string(vowelLow[k]) == string(str[i]) && k != len(vowelLow)-1 {
					vowIndex = k + 1
				} else if string(vowelLow[k]) == string(str[i]) && k == len(vowelLow)-1 {
					vowIndex = 0
				}
			}
			strConv += strings.ReplaceAll(string(str[i]), string(str[i]), string(vowelUp[vowIndex]))
		} else if strings.Contains(vowelUp, string(str[i])) == true {
			for k := 0; k < len(vowelUp); k++ {
				if string(vowelUp[k]) == string(str[i]) && k != len(vowelUp)-1 {
					vowIndex = k + 1
				} else if string(vowelUp[k]) == string(str[i]) && k == len(vowelUp)-1 {
					vowIndex = 0
				}
			}
			// fmt.Println(string(vowelUp[vowIndex]))
			strConv += strings.ReplaceAll(string(str[i]), string(str[i]), string(vowelLow[vowIndex]))
		} else {
			if unicode.IsLower(rune(str[i])) == true {
				strConv += strings.ReplaceAll(string(str[i]), string(str[i]), string(strings.ToUpper(string(str[i]))))
			} else if unicode.IsUpper(rune(str[i])) == true {
				strConv += strings.ReplaceAll(string(str[i]), string(str[i]), string(strings.ToLower(string(str[i]))))
			} else if string(str[i]) != " " {
				strConv += string(str[i])
			}
		}
	}

	return strConv // TODO: replace this
}

func CheckPassword(str string) string {
	strConv := Generate(str)
	char := "abcdefghijklmnopqrstuvwxyz1234567890"
	var strength string
	unChar := ""
	for i := 0; i < len(strConv); i++ {
		if strings.Contains(strings.ToUpper(char), string(strConv[i])) == false && strings.Contains(char, string(strConv[i])) == false {
			unChar += string(strConv[i])
		}
	}

	if len(strConv) >= 14 {
		for i := 0; i < len(strConv); i++ {
			if strings.Contains(char, string(strConv[i])) || strings.Contains(strings.ToUpper(char), string(strConv[i])) || strings.Contains(strings.ToUpper(char), string(strConv[i])) == false {
				strength = "kuat"
			}
		}
	} else if len(strConv) >= 7 {
		for i := 0; i < len(strConv); i++ {
			if strings.Contains(string(strConv[i]), unChar) == false {
				strength = "sedang"
			} else {
				strength = "lemah"
			}
		}
		// fmt.Println(unChar)
	} else {
		strength = "sangat lemah"
	}
	// fmt.Println(unChar)
	// fmt.Println(strength)
	return strength // TODO: replace this
}

func PasswordGenerator(base string) (string, string) {
	return Generate(base), CheckPassword(base) // TODO: replace this
}

func main() {
	data := "Semangat Pagi" // bisa digunakan untuk melakukan debug
	fmt.Println(PasswordGenerator(data))

	fmt.Println(PasswordGenerator("admin1238*(#@123i@DJHDJSUI@*(#)NDKJKJDSNJ"))
}
