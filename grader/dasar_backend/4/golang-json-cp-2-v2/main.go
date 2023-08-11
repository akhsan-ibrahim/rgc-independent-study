package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type LoanData struct {
	StartBalance int
	Data         []Loan
	Employees    []Employee
}

type Loan struct {
	Date        string
	EmployeeIDs []string
}

type Employee struct {
	ID       string
	Name     string
	Position string
}

// json structure
type LoanRecord struct {
	MonthDate    string     `json:"month_date"`
	StartBalance int        `json:"start_balance"`
	EndBalance   int        `json:"end_balance"`
	Borrowers    []Borrower `json:"borrowers"`
}

type Borrower struct {
	ID        string `json:"id"`
	TotalLoan int    `json:"total_loan"`
}

func LoanReport(data LoanData) LoanRecord {
	result := LoanRecord{}
	tempIDs := make(map[string]int)
	check := 0
	balance := data.StartBalance
	users := []Borrower{}
	tempUser := Borrower{}

	for _, val := range data.Data {
		spliter := strings.Split(val.Date, "-")
		result.MonthDate = spliter[1] + " " + spliter[2]

		for _, data := range val.EmployeeIDs {
			if balance >= 50000 {
				tempIDs[data] += 50000
				balance -= 50000
				check += 50000
			} else {
				tempIDs[data] += balance
				check += balance
			}
		}
	}
	// fmt.Println(tempIDs)

	for key, val := range tempIDs {
		// fmt.Println(key, val)
		tempUser.ID = key
		tempUser.TotalLoan = val
		users = append(users, tempUser)
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	result.EndBalance = balance
	result.StartBalance = data.StartBalance
	result.Borrowers = users

	return result
}

func RecordJSON(record LoanRecord, path string) error {
	jsonData, err := json.Marshal(record)
	if err != nil {
		return err
	}
	// fmt.Println(jsonData)
	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

// gunakan untuk debug
func main() {
	records := LoanReport(LoanData{
		StartBalance: 350000,
		Data: []Loan{
			{"01-March-2021", []string{"1", "2", "3", "4"}},
			{"04-March-2021", []string{"1", "2", "3"}},
		},
		Employees: []Employee{
			{"1", "Employee A", "Manager"},
			{"2", "Employee B", "Staff"},
			{"3", "Employee C", "Staff"},
			{"4", "Employee D", "Staff"},
		},
	})

	err := RecordJSON(records, "loan-records.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}
