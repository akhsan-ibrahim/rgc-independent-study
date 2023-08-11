package main

import "fmt"

func main() {
	// ini komentar
	// fmt.Println()
	// fmt.Println("Hello Go!")
	// fmt.Println()
	// fmt.Println("Hello Go!", 1234, true)
	// fmt.Printf("My name is %s\n", "Akhsan")
	// fmt.Printf("%d is my contact number\n", 8123)

	// var alamat, name string
	// fmt.Print("nama dan alamat: ")
	// fmt.Scan(&name, &alamat)
	// fmt.Println("Hello", name)
	// fmt.Println("Dari", alamat)

	// `
	// sdw
	// sda
	// `

	// fmt.Println("Golang"[0])
	// fmt.Println(string("Golang"[0]))

	// var a = 10
	// a %= 2
	// fmt.Println(a)

	// for i, total := 1, 0; i <= 5; i++ {
	// 	total += i
	// 	fmt.Println(total)
	// }

	// for i := 1; i < 5; i++ {
	// 	for j := 1; j < 5; j++ {
	// 		fmt.Printf("%d\t", i*j)
	// 	}
	// 	fmt.Println()
	// }

	// for i := 1; i < 5; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("* ")
	// 	}
	// 	fmt.Println()
	// }

	//i = 5
	//j = 5
	for i := 1; i < 5; i++ {
		for j := 5; j > i; j-- {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
