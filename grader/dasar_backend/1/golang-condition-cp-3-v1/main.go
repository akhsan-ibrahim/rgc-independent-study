package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	avg := (float64(math) + float64(science) + float64(english) + float64(indonesia)) / 4
	fmt.Println(avg)
	var result string

	if avg == 100 {
		result = "Sempurna"
	} else if avg >= 90 {
		result = "Sangat Baik"
	} else if avg >= 80 {
		result = "Baik"
	} else if avg >= 70 {
		result = "Cukup"
	} else if avg >= 60 {
		result = "Kurang"
	} else if avg < 60 {
		result = "Sangat kurang"
	}
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 50, 50, 60))
}
