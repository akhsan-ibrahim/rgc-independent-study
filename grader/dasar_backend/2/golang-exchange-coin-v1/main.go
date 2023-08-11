package main

import "fmt"

func ExchangeCoin(amount int) []int {
	digit := amount
	coins := []int{}
	for i := digit; i > 0; {
		if digit >= 1000 {
			digit = digit - 1000
			coins = append(coins, 1000)
			i -= 1000
		} else if digit >= 500 {
			digit = digit - 500
			coins = append(coins, 500)
			i -= 500
		} else if digit >= 200 {
			digit = digit - 200
			coins = append(coins, 200)
			i -= 200
		} else if digit >= 100 {
			digit = digit - 100
			coins = append(coins, 100)
			i -= 100
		} else if digit >= 50 {
			digit = digit - 50
			coins = append(coins, 50)
			i -= 50
		} else if digit >= 20 {
			digit = digit - 20
			coins = append(coins, 20)
			i -= 20
		} else if digit >= 10 {
			digit = digit - 10
			coins = append(coins, 10)
			i -= 10
		} else if digit >= 5 {
			digit = digit - 5
			coins = append(coins, 5)
			i -= 5
		} else if digit >= 1 {
			digit = digit - 1
			coins = append(coins, 1)
			i -= 1
		}
	}
	return coins // TODO: replace this
}

func main() {
	fmt.Println(ExchangeCoin(0))
}
