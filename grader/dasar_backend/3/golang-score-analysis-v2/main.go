package main

import (
	"fmt"
	"sort"
)

func Add(a, b int) int {
	return 0
}

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}
	sort.Ints(s.Grades)
	var sum, count float64
	for i, value := range s.Grades {
		sum += float64(value)
		count = float64(i) + 1
	}
	avg := sum / count
	min := s.Grades[0]
	max := s.Grades[len(s.Grades)-1]
	return avg, min, max // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	// avg, min, max := Analysis(School{
	// 	Name:    "Imam Assidiqi School",
	// 	Address: "Jl. Imam Assidiqi",
	// 	Grades:  []int{},
	// })

	// fmt.Println(avg, min, max)

	s := School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{},
	}
	s.AddGrade(100, 90, 100, 90, 100, 90)
	fmt.Println(s.Grades)
	fmt.Println(Analysis(s))
}
