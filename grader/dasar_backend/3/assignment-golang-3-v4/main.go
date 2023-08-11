package main

import (
	"a21hc3NpZ25tZW50/invoice"
	"fmt"
	"log"
)

func RecapDataInvoice(data []invoice.Invoice) ([]invoice.InvoiceData, error) {
	var result = []invoice.InvoiceData{}

	for _, inv := range data {
		// fmt.Println()
		val, err := inv.RecordInvoice()
		// fmt.Println(val)
		if err == nil {
			if len(result) != 0 {
				have := 0
				for i, departInv := range result {
					if departInv.Departemen == val.Departemen && departInv.Date == val.Date {
						departInv.TotalInvoice += val.TotalInvoice
						result[i] = departInv
						have = 1
					}
				}
				if have == 0 {
					result = append(result, val)
				}
			} else {
				result = append(result, val)
			}
		} else {
			return result, err
		}
	}
	// fmt.Println()
	return result, nil // TODO: replace this
}

func main() {
	listInvoice := []invoice.Invoice{
		invoice.FinanceInvoice{
			Date:     "01/02/2020",
			Details:  []invoice.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.FinanceInvoice{
			Date:     "02/02/2020",
			Details:  []invoice.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
			Status:   invoice.PAID,
			Approved: true,
		},
		invoice.WarehouseInvoice{
			Date: "03-February-2020",
			Products: []invoice.Product{
				{"product A", 10, 10000, 0.1},
				{"product C", 5, 15000, 0.2},
			},
			InvoiceType: invoice.PURCHASE,
			Approved:    true,
		},
		invoice.MarketingInvoice{
			Date:        "04/02/2020",
			StartDate:   "20/01/2020",
			EndDate:     "25/01/2020",
			Approved:    true,
			PricePerDay: 10000,
			AnotherFee:  5000,
		},
	}

	result, err := RecapDataInvoice(listInvoice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	// fmt.Println(listInvoice)
}
