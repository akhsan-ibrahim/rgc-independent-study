package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
	var set []int
	if len(villager) == 0 {
		return []int{}
	} else {
		set = villager[0]
	}

	for i := 0; i < len(villager); i++ {
		var comp []int
		for j := 0; j < len(villager[i]); j++ {
			for k := 0; k < len(set); k++ {
				if villager[i][j] == set[k] {
					comp = append(comp, villager[i][j])
				}
			}
		}
		set = comp
		fmt.Println(comp)
	}
	return set // TODO: replace this
}
func main() {
	villager := [][]int{
		{7, 12, 19, 22},
		{12, 19, 21, 23},
		{7, 12, 19},
		{12, 19},
	}
	// villager := [][]int{}
	fmt.Println(SchedulableDays(villager))
}
