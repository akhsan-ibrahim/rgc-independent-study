package invoice

import (
	"errors"
	"strconv"
	"strings"
)

// import "errors"

// Marketing invoice
type MarketingInvoice struct {
	Date        string
	StartDate   string
	EndDate     string
	PricePerDay int
	AnotherFee  int
	Approved    bool
}

func (mi MarketingInvoice) RecordInvoice() (InvoiceData, error) {
	if mi.Date == "" {
		return InvoiceData{}, errors.New("invoice date is empty")
	}
	if mi.StartDate == "" || mi.EndDate == "" {
		return InvoiceData{}, errors.New("travel date is empty")
	}
	if mi.PricePerDay <= 0 || mi.AnotherFee <= 0 {
		return InvoiceData{}, errors.New("price per day is not valid")
	}
	date := ChangeDate(mi.Date)
	slcDateStr := strings.Split(mi.StartDate, "/")
	slcDateEnd := strings.Split(mi.EndDate, "/")
	dateStr, _ := strconv.ParseFloat(slcDateStr[0], 64)
	dateEnd, _ := strconv.ParseFloat(slcDateEnd[0], 64)
	total := (dateEnd-dateStr+1)*float64(mi.PricePerDay) + float64(mi.AnotherFee)
	var result = InvoiceData{
		Date:         date,
		TotalInvoice: total,
		Departemen:   Marketing,
	}
	return result, nil // TODO: replace this
}
