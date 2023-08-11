package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	var result []int
	if len(data) == 0 {
		return []int{}
	} else {
		for _, branch := range data {
			for j, month := range branch {
				profit := month[0] - month[1]
				if len(result) == j {
					result = append(result, profit)
				} else {
					result[j] += profit
				}
			}
		}
	}
	return result // TODO: replace this
}
func main() {
	data := [][][2]int{{{1000, 800}, {700, 500}}, {{1000, 800}, {900, 200}}}
	// data := [][][2]int{}
	fmt.Println(CountProfit(data))
}
