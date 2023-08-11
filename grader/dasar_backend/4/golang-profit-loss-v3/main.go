package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	var content, _ = os.ReadFile(path)
	// fmt.Println(string(content))
	data := string(content)
	// fmt.Println(len(data))
	if len(data) == 0 {
		return []string{}, nil
	}
	var lsData = strings.Split(data, "\n")
	// fmt.Println(len(lsData))
	return lsData, nil // TODO: replace this
}

func CalculateProfitLoss(data []string) string {
	var date, status string
	result := 0
	for i := 0; i < len(data); i++ {
		aData := strings.Split(data[i], ";")
		amount, _ := strconv.Atoi(aData[2])
		if aData[1] == "income" {
			result += amount
		} else if aData[1] == "expense" {
			result -= amount
		}
		date = aData[0]
	}
	if result > 0 {
		status = "profit"
	} else {
		status = "loss"
	}
	result = int(math.Abs(float64(result)))
	strResult := strconv.Itoa(result)
	return date + ";" + status + ";" + strResult // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	// result := CalculateProfitLoss(datas)
	// fmt.Println(result)

	fmt.Println(datas)
}
