package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	var result string
	for i := 0; i < len(data); i++ {
		if strings.Contains(data[i], input) == true {
			result += data[i] + ","
		}
	}
	return result[:len(result)-1] // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
