package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	noChar := "!@#$%^&*()-+=.,),"
	var result string
	strSliceWord := strings.Split(str, " ") //nyeplit per kata

	for i := 0; i < len(strSliceWord); i++ {
		var strReverse []string                              //nyimpen reverse kalimat
		strSliceLetter := strings.Split(strSliceWord[i], "") //nyeplit huruf di setiap kata
		for j := len(strSliceLetter) - 1; j >= 0; j-- {
			if strings.Contains(noChar, strSliceLetter[j]) == false {
				strReverse = append(strReverse, strSliceLetter[j]) //menambahkan hasil reverse ke var final
			}
		}

		if len(strReverse) == 0 {
			return ""
			break
		} else {
			lastLetter := []rune(strReverse[len(strReverse)-1]) //menyimpan karakter terakhir dari slice final
			if unicode.IsUpper(lastLetter[0]) {
				strReverse[len(strReverse)-1] = strings.ToLower(strReverse[len(strReverse)-1])
				strReverse[0] = strings.ToUpper(strReverse[0])
			}
		}

		for k := 0; k < len(strReverse); k++ {
			word := []rune(strReverse[k])
			if unicode.IsUpper(word[0]) && k != 0 {
				result += strings.ToLower(strReverse[k]) //mengecilkan huruf selain index 0 yg kapital
			} else {
				result += strReverse[k]
			}
		}
		result += " "
	}

	return result[:len(result)-1] // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
	fmt.Println(ReverseWord(""))
}
