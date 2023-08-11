package invoice

import "errors"

// Warehouse invoice

type WarehouseInvoice struct {
	Date        string
	InvoiceType InvoiceTypeName
	Approved    bool
	Products    []Product
}

type InvoiceTypeName string

const (
	PURCHASE InvoiceTypeName = "purchase"
	SALES    InvoiceTypeName = "sales"
)

type Product struct {
	Name     string
	Unit     int
	Price    float64
	Discount float64
}

func (wi WarehouseInvoice) RecordInvoice() (InvoiceData, error) {
	if wi.Date == "" {
		return InvoiceData{}, errors.New("invoice date is empty")
	}
	if wi.InvoiceType != PURCHASE && wi.InvoiceType != SALES {
		return InvoiceData{}, errors.New("invoice type is not valid")
	}
	if len(wi.Products) == 0 {
		return InvoiceData{}, errors.New("invoice products is empty")
	}

	var total float64
	for _, val := range wi.Products {
		if val.Unit <= 0 {
			return InvoiceData{}, errors.New("unit product is not valid")
		}
		if val.Price <= 0 {
			return InvoiceData{}, errors.New("price product is not valid")
		}
		prices := float64(val.Unit) * val.Price
		total += prices - (val.Discount * prices)
	}
	var result = InvoiceData{
		Date:         wi.Date,
		TotalInvoice: total,
		Departemen:   Warehouse,
	}
	return result, nil // TODO: replace this
}
