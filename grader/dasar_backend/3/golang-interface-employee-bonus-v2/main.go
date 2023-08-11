package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

func (j Junior) GetBonus() float64 {
	var prorata float64
	if j.WorkingMonth > 12 {
		prorata = 1
	} else {
		prorata = float64(j.WorkingMonth) / 12
	}
	return float64(j.BaseSalary) * prorata
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

func (s Senior) GetBonus() float64 {
	var prorata float64
	if s.WorkingMonth > 12 {
		prorata = 1
	} else {
		prorata = float64(s.WorkingMonth) / 12
	}
	return 2*float64(s.BaseSalary)*prorata + (s.PerformanceRate * float64(s.BaseSalary))
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (m Manager) GetBonus() float64 {
	var prorata float64
	if m.WorkingMonth > 12 {
		prorata = 1
	} else {
		prorata = float64(m.WorkingMonth) / 12
	}
	return 2*float64(m.BaseSalary)*prorata + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus() // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	var total float64
	for _, emp := range employees {
		// fmt.Println(emp)
		total += emp.GetBonus()
	}
	return total // TODO: replace this
}

func main() {
	// var emp Employee
	// emp = Junior{Name: "Aku", BaseSalary: 100000, WorkingMonth: 12}
	// emp = Senior{Name: "Aku", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}
	// emp = Manager{Name: "Aku", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5}
	// fmt.Println(EmployeeBonus(emp))

	var emps = []Employee{
		Junior{Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12},
		Senior{Name: "Senior A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5},
		Manager{Name: "Manager A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5},
	}
	fmt.Println(TotalEmployeeBonus(emps))
}
