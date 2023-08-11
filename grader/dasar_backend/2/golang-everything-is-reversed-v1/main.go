package main

import (
	"fmt"
	"strconv"
)

func ReverseData(arr [5]int) [5]int {
	var revs [5]int
	var revsNumber [5]int
	for i := 0; i < len(arr); i++ {
		var revNumber string
		for j := len(strconv.Itoa(arr[i])) - 1; j >= 0; j-- {
			revNumber += string(strconv.Itoa(arr[i])[j])
		}
		intRevNumber, _ := strconv.Atoi(revNumber)
		revs[i] = intRevNumber
	}
	revsNumber[0] = revs[4]
	revsNumber[1] = revs[3]
	revsNumber[2] = revs[2]
	revsNumber[3] = revs[1]
	revsNumber[4] = revs[0]
	return revsNumber // TODO: replace this
}

func main() {
	var arr = [5]int{456789, 44332, 2221, 12, 10}
	fmt.Println(ReverseData(arr))
}
