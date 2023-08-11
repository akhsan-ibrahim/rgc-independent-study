package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	strNumb := strconv.Itoa(numbers)
	var lsFinalNum []int
	var lsSum []int

	for i := 0; i < len(strNumb)-1; i++ {
		numb1, _ := strconv.Atoi(string(strNumb[i]))
		numb2, _ := strconv.Atoi(string(strNumb[i+1]))
		var sum int
		var finalNum int
		if (numb1 >= 0 && numb1 <= 9) && (len(strNumb) >= 2 && len(strNumb) <= 100000) {
			sum = numb1 + numb2
			lsSum = append(lsSum, sum)
			strNum := strconv.Itoa(numb1) + strconv.Itoa(numb2)
			finalNum, _ = strconv.Atoi(string(strNum))
			lsFinalNum = append(lsFinalNum, finalNum)
		} else {
			return 0
			break
		}
	}

	var biggest int
	var pair int
	for i := 0; i < len(lsSum); i++ {
		if lsSum[i] > biggest {
			biggest = lsSum[i]
			pair = lsFinalNum[i]
		}
	}
	return pair // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(89083278))
	fmt.Println(BiggestPairNumber(11223344))
}
