package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	var result [][]string
	for key, value := range mapData {
		var data []string
		// fmt.Println(key + "\t: " + value)
		data = append(data, key, value)
		result = append(result, data)
	}
	// fmt.Println()
	return result // TODO: replace this
}

func main() {
	data := map[string]string{
		"hello": "world",
		"John":  "Doe",
		"age":   "14",
	}

	fmt.Println(MapToSlice(data))
}
