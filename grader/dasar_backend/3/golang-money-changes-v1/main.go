package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	var coin = []int{}
	var total int
	for _, item := range products {
		// fmt.Println(item)
		total += item.Price + item.Tax
	}
	change := amount - total

	for i := change; i > 0; i-- {
		if change >= 1000 {
			coin = append(coin, 1000)
			change -= 1000
			i = change
		} else if change >= 500 {
			coin = append(coin, 500)
			change -= 500
			i = change
		} else if change >= 200 {
			coin = append(coin, 200)
			change -= 200
			i = change
		} else if change >= 100 {
			coin = append(coin, 100)
			change -= 100
			i = change
		} else if change >= 50 {
			coin = append(coin, 50)
			change -= 50
			i = change
		} else if change >= 20 {
			coin = append(coin, 20)
			change -= 20
			i = change
		} else if change >= 10 {
			coin = append(coin, 10)
			change -= 10
			i = change
		} else if change >= 5 {
			coin = append(coin, 5)
			change -= 5
			i = change
		} else if change >= 1 {
			coin = append(coin, 1)
			change -= 1
			i = change
		}
	}

	// fmt.Println(amount, total, change)
	return coin // TODO: replace this
}

func main() {
	buy := []Product{
		{Name: "Baju", Price: 5000, Tax: 500},
		{Name: "Celana", Price: 3000, Tax: 300},
	}
	fmt.Println(MoneyChanges(10000, buy))
}
