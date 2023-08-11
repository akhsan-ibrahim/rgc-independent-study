package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverseWord(str string) string {
	spcl := "!@#$%^&*()-+=.,),"
	result := ""

	strSlice := strings.Split(str, " ") //nyeplit per kata
	for i := 0; i < len(strSlice); i++ {
		var final []string                            //nyimpen reverse kalimat
		strSliceNew := strings.Split(strSlice[i], "") //nyeplit huruf di setiap kata
		for j := len(strSliceNew) - 1; j >= 0; j-- {
			if strings.Contains(spcl, strSliceNew[j]) == false {
				final = append(final, strSliceNew[j]) //menambahkan hasil reverse ke var final
			}
		}

		r := []rune(final[len(final)-1]) //menyimpan karakter terakhir dari slice final
		if unicode.IsUpper(r[0]) {
			final[len(final)-1] = strings.ToLower(final[len(final)-1])
			final[0] = strings.ToUpper(final[0])
		}
		fmt.Println(final)
		for k := 0; k < len(final); k++ {
			r := []rune(final[k])
			if unicode.IsUpper(r[0]) && k != 0 {
				result += strings.ToLower(final[k]) //mengecilkan huruf selain index 0 yg kapital
			} else {
				result += final[k]
			}

		}
		result += " "
	}

	return result
}

func main() {
	fmt.Println(reverseWord("Aku Sayang Ibu"))
	fmt.Println(reverseWord("A bird fly to the Sky"))
	fmt.Println(reverseWord("ini terlalu mudah"))
	fmt.Println(reverseWord(("KITA SELALU BERSAMA")))
}
