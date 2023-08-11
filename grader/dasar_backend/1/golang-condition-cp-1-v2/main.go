package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	var hasil string

	if score >= 0 && score <= 100 {
		if absent >= 0 && absent <= 10 {
			if score >= 70 && absent < 5 {
				hasil = "lulus"
			} else if score < 70 || absent >= 5 {
				hasil = "tidak lulus"
			}
		} else {
			hasil = "Masukkan absen 0-10"
		}
	} else {
		hasil = "Masukkan nilai 0-100"
	}
	return hasil // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(100, 4))
	fmt.Println(GraduateStudent(80, 5))
}
