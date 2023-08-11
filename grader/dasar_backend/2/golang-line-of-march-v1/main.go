package main

import "fmt"

func Sortheight(height []int) []int {
	var sortHeight = height
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height)-i-1; j++ {
			if height[j] > height[j+1] {
				poin := height[j]
				height[j] = height[j+1]
				height[j+1] = poin
			}
		}
	}
	return sortHeight // TODO: replace this
}

func main() {
	var height = []int{172, 170, 150, 165, 144, 155, 159}
	fmt.Println(Sortheight(height))
}
