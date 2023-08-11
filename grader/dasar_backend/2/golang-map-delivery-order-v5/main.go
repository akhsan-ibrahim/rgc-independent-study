package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here

func DeliveryOrder(data []string, day string) map[string]float32 {
	dayLocAdm := make(map[string]map[string]float32)
	dayLocAdm["senin"] = map[string]float32{
		"JKT": 0.1,
		"DPK": 0.1,
	}
	dayLocAdm["selasa"] = map[string]float32{
		"JKT": 0.05,
		"BKS": 0.05,
		"DPK": 0.05,
	}
	dayLocAdm["rabu"] = map[string]float32{
		"JKT": 0.1,
		"BDG": 0.1,
	}
	dayLocAdm["kamis"] = map[string]float32{
		"JKT": 0.05,
		"BKS": 0.05,
		"BDG": 0.05,
	}
	dayLocAdm["jumat"] = map[string]float32{
		"JKT": 0.1,
		"BKS": 0.1,
	}
	dayLocAdm["sabtu"] = map[string]float32{
		"JKT": 0.05,
		"BDG": 0.05,
	}

	result := make(map[string]float32)
	for _, desc := range data {
		descSplit := strings.Split(desc, ":")
		fullName := descSplit[0] + "-" + descSplit[1]
		cost, _ := strconv.ParseFloat(descSplit[2], 32)
		location := descSplit[3]
		if dayLocAdm[day][location] != 0 {
			result[fullName] = float32(cost) + (float32(cost) * dayLocAdm[day][location])
		} else {
			continue
		}
	}
	return result // TODO: replace this
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
