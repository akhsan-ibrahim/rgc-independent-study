package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	var result = []int{}
	for i := 0; i < len(date1); i++ {
		for j := 0; j < len(date2); j++ {
			if date1[i] == date2[j] {
				result = append(result, date1[i])
			}
		}
	}
	return result // TODO: replace this
}

func main() {
	var date1 = []int{11, 12, 13, 14, 15}
	var date2 = []int{5, 10, 12, 13, 20, 21}
	fmt.Println(SchedulableDays(date1, date2))
}
