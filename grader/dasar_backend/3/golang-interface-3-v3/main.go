package main

import (
	"fmt"
	"strconv"
	// "golang.org/x/tools/go/analysis/passes/printf"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time any) string {
	var result string
	if _, ok := time.(string); ok {
		if len(time.(string)) != 5 || string(time.(string)[2]) != ":" {
			return "Invalid input"
		}
		hour, _ := strconv.Atoi(time.(string)[:2])
		strHour := fmt.Sprintf("%02d", hour)
		if hour < 12 {
			result = strHour + time.(string)[2:] + " AM"
		} else if hour == 12 {
			result = strHour + time.(string)[2:] + " PM"
		} else {
			hour -= 12
			strHour := fmt.Sprintf("%02d", hour)
			result = strHour + time.(string)[2:] + " PM"
		}
	} else if _, ok := time.([]int); ok {
		if len(time.([]int)) != 2 {
			return "Invalid input"
		}
		hour := time.([]int)[0]
		strHour := fmt.Sprintf("%02d", hour)
		strMinute := fmt.Sprintf("%02d", time.([]int)[1])
		if hour < 12 {
			result = strHour + ":" + strMinute + " AM"
		} else if hour == 12 {
			result = strHour + ":" + strMinute + " PM"
		} else {
			hour -= 12
			strHour = fmt.Sprintf("%02d", hour)
			result = strHour + ":" + strMinute + " PM"
		}
	} else if _, ok := time.(map[string]int); ok {
		// fmt.Println(time)
		mapTime := time.(map[string]int)
		hour, isHour := mapTime["hour"]
		minute, isMinute := mapTime["minute"]
		if !isHour || !isMinute {
			return "Invalid input"
		}
		strHour := fmt.Sprintf("%02d", hour)
		strMinute := fmt.Sprintf("%02d", minute)
		if hour < 12 {
			result = strHour + ":" + strMinute + " AM"
		} else if hour == 12 {
			result = strHour + ":" + strMinute + " PM"
		} else {
			hour -= 12
			strHour = fmt.Sprintf("%02d", hour)
			result = strHour + ":" + strMinute + " PM"
		}
	} else if _, ok := time.(Time); ok {
		// fmt.Println(time)
		strucTime := time.(Time)
		hour := strucTime.Hour
		minute := strucTime.Minute
		strHour := fmt.Sprintf("%02d", hour)
		strMinute := fmt.Sprintf("%02d", minute)
		if hour < 12 {
			result = strHour + ":" + strMinute + " AM"
		} else if hour == 12 {
			result = strHour + ":" + strMinute + " PM"
		} else {
			hour -= 12
			strHour = fmt.Sprintf("%02d", hour)
			result = strHour + ":" + strMinute + " PM"
		}
	} else {
		return "Invalid input"
	}
	return result // TODO: replace this
}

func main() {
	fmt.Println(ChangeToStandartTime("04:00"))
	fmt.Println(ChangeToStandartTime([]int{14, 00}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 02, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
