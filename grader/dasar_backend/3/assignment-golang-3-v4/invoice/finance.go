package invoice

import "errors"

// Finance invoice
type FinanceInvoice struct {
	Date     string
	Status   InvoiceStatus // status: "paid", "unpaid"
	Approved bool
	Details  []Detail
}

type InvoiceStatus string

const (
	PAID   InvoiceStatus = "paid"
	UNPAID InvoiceStatus = "unpaid"
)

type Detail struct {
	Description string
	Total       int
}

func (fi FinanceInvoice) RecordInvoice() (InvoiceData, error) {
	if fi.Date == "" {
		return InvoiceData{}, errors.New("invoice date is empty")
	}
	if fi.Status != PAID && fi.Status != UNPAID {
		return InvoiceData{}, errors.New("invoice status is not valid")
	}
	if len(fi.Details) == 0 {
		return InvoiceData{}, errors.New("invoice details is empty")
	}

	if fi.Status == UNPAID || fi.Approved == false {
		return InvoiceData{}, errors.New("status unpaid/unapproved")
	}

	date := ChangeDate(fi.Date)
	var total float64
	for _, val := range fi.Details {
		total += float64(val.Total)
	}

	if total <= 0 {
		return InvoiceData{}, errors.New("total price is not valid")
	}

	var record = InvoiceData{
		Date:         date,
		TotalInvoice: total,
		Departemen:   Finance,
	}
	return record, nil // TODO: replace this
}
