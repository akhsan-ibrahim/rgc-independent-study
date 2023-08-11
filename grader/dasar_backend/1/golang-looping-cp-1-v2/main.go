package main

import "fmt"

func CountingNumber(n int) float64 {
	var i, result float64

	for i = 1; i <= float64(n); i += 0.5 {
		result += i
		// fmt.Println(result)
	}
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
