package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var hasil float64

	if height >= 130 && height <= 250 {
		if gender == "laki-laki" {
			hasil = float64(height-100) - (float64(height-100) * 0.1)
		} else if gender == "perempuan" {
			hasil = float64(height-100) - (float64(height-100) * 0.15)
		} else {
			fmt.Println("Masukkan gender laki-laki/perempuan")
		}
	} else {
		fmt.Println("Masukkan tinggi 130-250")
	}

	return hasil // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	// fmt.Println(BMICalculator("laki-laki", 155))
	fmt.Println(BMICalculator("perempuan", 155))
}
