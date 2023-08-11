package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {
	var strDay string
	var strMonth string
	strYear := strconv.Itoa(year)
	var maxDay int
	switch month {
	case 1:
		strMonth = "January"
		maxDay = 31
	case 2:
		strMonth = "February"
		maxDay = 29
	case 3:
		strMonth = "March"
		maxDay = 31
	case 4:
		strMonth = "April"
		maxDay = 30
	case 5:
		strMonth = "May"
		maxDay = 31
	case 6:
		strMonth = "June"
		maxDay = 30
	case 7:
		strMonth = "July"
		maxDay = 31
	case 8:
		strMonth = "August"
		maxDay = 31
	case 9:
		strMonth = "September"
		maxDay = 30
	case 10:
		strMonth = "October"
		maxDay = 31
	case 11:
		strMonth = "November"
		maxDay = 30
	case 12:
		strMonth = "December"
		maxDay = 31
	default:
		return "Invalid month"
		break
	}

	if day > 0 && day <= maxDay {
		if day < 10 {
			strDay = "0" + strconv.Itoa(day)
		} else {
			strDay = strconv.Itoa(day)
		}
	} else {
		return "Invalid day"
	}

	result := strDay + "-" + strMonth + "-" + strYear
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
