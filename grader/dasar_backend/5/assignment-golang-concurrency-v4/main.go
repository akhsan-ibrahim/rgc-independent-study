package main

import (
	"fmt"
	// "time"
)

type RowData struct {
	RankWebsite int
	Domain      string
	TLD         string
	IDN_TLD     string
	Valid       bool
	RefIPs      int
}

func GetTLD(domain string) (TLD string, IDN_TLD string) {
	var ListIDN_TLD = map[string]string{
		".com": ".co.id",
		".org": ".org.id",
		".gov": ".go.id",
	}

	for i := len(domain) - 1; i >= 0; i-- {
		if domain[i] == '.' {
			TLD = domain[i:]
			break
		}
	}

	if _, ok := ListIDN_TLD[TLD]; ok {
		return TLD, ListIDN_TLD[TLD]
	} else {
		return TLD, TLD
	}
}

func ProcessGetTLD(website RowData, ch chan RowData, chErr chan error) {
	if website.Domain == "" {
		chErr <- fmt.Errorf("domain name is empty")
		return
	}
	if website.Valid == false {
		chErr <- fmt.Errorf("domain not valid")
		return
	}
	if website.RefIPs == -1 {
		chErr <- fmt.Errorf("domain RefIPs not valid")
		return
	}

	TLD, IDN_TLD := GetTLD(website.Domain)
	website.TLD = TLD
	website.IDN_TLD = IDN_TLD
	ch <- website
	// chErr <- nil
}

// Gunakan variable ini sebagai goroutine di fungsi FilterAndGetDomain
var FuncProcessGetTLD = ProcessGetTLD

func FilterAndFillData(TLD string, data []RowData) ([]RowData, error) {
	ch := make(chan RowData, len(data))
	errCh := make(chan error)

	for _, website := range data {
		go FuncProcessGetTLD(website, ch, errCh)
	}

	var chLs []RowData

	for i := 0; i < len(data); i++ {
		select {
		case err := <-errCh:
			return chLs, err
		default:
			item := <-ch
			if item.TLD == TLD {
				chLs = append(chLs, item)
			}
		}
	}
	return chLs, nil // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	rows, err := FilterAndFillData(".com", []RowData{
		{1, "", "", "", true, 100},
		{2, "facebook.com", "", "", true, 100},
		{3, "golang.org", "", "", true, 100},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)

	// webs, _ :=
}
