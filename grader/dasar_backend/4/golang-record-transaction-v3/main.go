package main

import (
	// "errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	transactions = append(transactions, Transaction{"end", "end", -1})
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date < transactions[j].Date
	})

	baseDate := transactions[0].Date
	var transType string
	var transAmount int
	var result = []string{}
	for i := 0; i < len(transactions)-1; i++ {
		if transactions[i].Type == "income" {
			transAmount += transactions[i].Amount
			// fmt.Println("in:", transAmount)
		} else if transactions[i].Type == "expense" {
			transAmount -= transactions[i].Amount
			// fmt.Println("ex:", transAmount)
		}

		if transAmount > 0 {
			transType = "income"
		} else {
			transType = "expense"
		}

		if transactions[i+1].Date != baseDate {
			result = append(result, baseDate+";"+transType+";"+strconv.Itoa(int(math.Abs(float64(transAmount)))))
			baseDate = transactions[i+1].Date
			transAmount = 0
		}
	}

	content := strings.Join(result, "\n")
	err := os.WriteFile(path, []byte(content), 0644)

	return err // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"03/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"02/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
		// fmt.Println(err)
	}

	fmt.Println("Success")
}
