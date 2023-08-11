package main

import "fmt"

func TicketPlayground(height, age int) int {
	var ticket int

	if age > 12 {
		ticket = 100_000
	} else if age == 12 || height > 160 {
		ticket = 60_000
	} else if age >= 10 && age <= 11 || height > 150 {
		ticket = 40_000
	} else if age >= 8 && age <= 9 || height > 135 {
		ticket = 25_000
	} else if age >= 5 && age <= 7 || height > 120 {
		ticket = 15_000
	} else {
		ticket = -1
	}

	return ticket // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
