package invoice

import (
	"strconv"
	"strings"
)

type InvoiceData struct {
	Date         string
	TotalInvoice float64
	Departemen   DepartmentName
}

type DepartmentName string

const (
	Finance   DepartmentName = "finance"
	Warehouse DepartmentName = "warehouse"
	Marketing DepartmentName = "marketing"
)

type Invoice interface {
	RecordInvoice() (InvoiceData, error)
}

// for utility
// change date 01/01/2022 -> 01-january-2022
func ChangeDate(date string) string {
	var arrMonth = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	splDate := strings.Split(date, "/")

	month, _ := strconv.Atoi(splDate[1])
	month = month - 1

	return splDate[0] + "-" + arrMonth[month] + "-" + splDate[2]
}
