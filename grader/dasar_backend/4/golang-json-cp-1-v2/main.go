package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "golang.org/x/text/date"
)

type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
	// TODO: answer here
}

type Study struct {
	Study_name   string `json:"study_name"`
	Study_credit int    `json:"study_credit"`
	Grade        string `json:"grade"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	content, _ := ioutil.ReadFile(filename)

	var report Report
	err := json.Unmarshal(content, &report)
	if err != nil {
		return Report{}, err
	}

	return report, nil // TODO: answer here
}

func GradePoint(report Report) float64 {
	var total_grade, total_studyCredit float64
	for _, study := range report.Studies {
		// fmt.Println(study)
		var scale float64
		letter := study.Grade
		switch letter {
		case "A":
			scale = 4
		case "AB":
			scale = 3.5
		case "B":
			scale = 3
		case "BC":
			scale = 2.5
		case "C":
			scale = 2
		case "CD":
			scale = 1.5
		case "D":
			scale = 1
		case "DE":
			scale = 0.5
		case "E":
			scale = 0
		}

		total_grade += scale * float64(study.Study_credit)
		total_studyCredit += float64(study.Study_credit)
	}

	var result float64
	if len(report.Studies) == 0 {
		return result
	}
	result = total_grade / total_studyCredit
	// fmt.Println(result)
	return result // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
	// fmt.Println(report)
}
