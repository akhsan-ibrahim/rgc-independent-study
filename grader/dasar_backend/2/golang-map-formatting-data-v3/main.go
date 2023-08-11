package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	itemMap := make(map[string]map[int][][]string)
	result := make(map[string][]string)
	for _, str := range data {
		strSplit := strings.Split(str, "-")
		var pos [][]string
		if _, ok := itemMap[strSplit[0]]; !ok {
			ind, _ := strconv.Atoi(strSplit[1])
			if _, ok := itemMap[strSplit[0]][ind]; !ok {
				pos = append(pos, strSplit[2:])
				itemMap[strSplit[0]] = map[int][][]string{
					ind: pos,
				}
			}
		} else {
			ind, _ := strconv.Atoi(strSplit[1])
			if _, ok := itemMap[strSplit[0]][ind]; !ok {
				pos = append(pos, strSplit[2:])
				itemMap[strSplit[0]][ind] = pos
			} else {
				pos = append(pos, strSplit[2:])
				itemMap[strSplit[0]][ind] = append(itemMap[strSplit[0]][ind], pos...)
			}
		}
	}
	// fmt.Println(itemMap)

	for key, value := range itemMap {
		// fmt.Println()
		// fmt.Println(key+" :", value)
		var fullName []string
		for i := 0; i < len(value); i++ {
			pos := value[i]
			// fmt.Println(i, pos)
			for _, val := range pos {
				// fmt.Println(val)
				if val[0] == "first" {
					fullName = append(fullName, val[1])
					result[key] = fullName
				} else {
					// fullName = append(fullName, val[1])
					fullName[i] = fullName[i] + " " + val[1]
					result[key] = fullName
				}
			}
		}
	}
	// fmt.Println()cd ..
	return result // TODO: replace this
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
