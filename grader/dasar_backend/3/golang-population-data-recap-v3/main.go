package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]any {
	var result = []map[string]any{}
	for _, person := range data {
		persona := strings.Split(person, ";")
		dataMap := make(map[string]any)
		dataMap["name"] = persona[0]
		age, _ := strconv.Atoi(persona[1])
		dataMap["age"] = age
		dataMap["address"] = persona[2]
		if len(persona[3]) != 0 {
			height, _ := strconv.ParseFloat(persona[3], 2)
			dataMap["height"] = height
		}
		if len(persona[4]) != 0 {
			married, _ := strconv.ParseBool(persona[4])
			dataMap["isMarried"] = married
		}
		// fmt.Println(dataMap)
		result = append(result, dataMap)
	}
	// fmt.Println()
	return result // TODO: replace this
}

func main() {
	data := []string{
		"Budi;23;Jakarta;;",
		"Joko;30;Bandung;;true",
		"Susi;25;Bogor;165.42;",
	}
	fmt.Println(PopulationData(data))
}
