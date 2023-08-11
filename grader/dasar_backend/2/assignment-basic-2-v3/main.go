package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi ini akan mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
func FormatRupiah(number int) string {
	var result string

	numStr := strconv.Itoa(number)

	for i := len(numStr) - 1; i >= 0; i-- {
		result = string(numStr[i]) + result
		if (len(numStr)-i)%3 == 0 && i != 0 {
			result = "." + result
		}
	}

	return "Rp " + result
}

func GetDayDifference(date string) int {
	var dateOfMonth = map[string][]int{
		"January":   []int{1, 30},
		"February":  []int{2, 28},
		"March":     []int{3, 31},
		"April":     []int{4, 30},
		"May":       []int{5, 31},
		"June":      []int{6, 30},
		"July":      []int{7, 31},
		"August":    []int{8, 31},
		"September": []int{9, 30},
		"October":   []int{10, 31},
		"November":  []int{11, 30},
		"December":  []int{12, 31},
	}

	dateSlc := strings.Split(date, " ")
	if len(dateSlc) > 6 {
		return -1
	}

	year, _ := strconv.Atoi(dateSlc[5])
	if year%4 == 0 {
		dateOfMonth["February"][1] = 29
	}

	var diff int
	firstDate, _ := strconv.Atoi(dateSlc[0])
	lastDate, _ := strconv.Atoi(dateSlc[3])
	if dateSlc[1] == dateSlc[4] {
		diff = lastDate - firstDate + 1
	} else {
		diff = (dateOfMonth[dateSlc[1]][1] - firstDate + 1) + lastDate

		for i := dateOfMonth[dateSlc[1]][0] + 1; i < dateOfMonth[dateSlc[4]][0]; i++ {
			// fmt.Println(i)
			for month, value := range dateOfMonth {
				if value[0] == i {
					// fmt.Println(month)
					diff += dateOfMonth[month][1] - 1
				}
			}
		}

	}
	// fmt.Println(diff)
	return diff // TODO: replace this
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	var attend = map[string]int{}
	if rangeDay > len(data) {
		rangeDay = len(data)
	}
	for i := 0; i < rangeDay; i++ {
		// fmt.Println(data[i])
		for _, name := range data[i] {
			if _, ok := attend[name]; !ok {
				attend[name] = 50000
			} else {
				attend[name] += 50000
			}
		}
	}
	var salary = map[string]string{}
	for key, _ := range attend {
		salary[key] = FormatRupiah(attend[key])
	}
	return salary // TODO: replace this
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	diff := GetDayDifference(dateRange)

	return GetSalary(diff, data) // TODO: replace this
}

func main() {
	res := GetSalaryOverview("30 March - 2 May 2021", [][]string{
		// {"andi", "Rojaki", "raji", "supri"},
		// {"andi", "Rojaki", "raji"},
		// {"andi", "raji", "supri"},
		// {"andi", "Rojaki", "raji", "supri"},
		{"A", "B"},
	})

	fmt.Println(res)

	// fmt.Println(GetDayDifference("25 November - 10 December 2021"))
	// fmt.Println(GetSalary(4, [][]string{
	// 	{"andi", "Rojaki", "raji", "supri"},
	// 	{"andi", "Rojaki", "raji"},
	// 	{"andi", "raji", "supri"},
	// 	{"andi", "Rojaki", "raji", "supri"},
	// }))
}
