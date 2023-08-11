package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	calc := strings.Split(calculate, " ")
	first, _ := strconv.ParseFloat(calc[0], 32)

	c := internal.Calculator{Base: float32(first)}

	for i := 1; i < len(calc); i++ {
		if calc[i] == "+" {
			num, _ := strconv.ParseFloat(calc[i+1], 32)
			c.Add(float32(num))
		} else if calc[i] == "-" {
			num, _ := strconv.ParseFloat(calc[i+1], 32)
			c.Subtract(float32(num))
		} else if calc[i] == "*" {
			num, _ := strconv.ParseFloat(calc[i+1], 32)
			c.Multiply(float32(num))
		} else if calc[i] == "/" {
			num, _ := strconv.ParseFloat(calc[i+1], 32)
			c.Divide(float32(num))
		} else {
			continue
		}
	}
	// fmt.Println(c)
	return c.Result() // TODO: replace this
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
