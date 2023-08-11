package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	var poinDomain, poinTLD, last int
	var domain, TLD string

	poinDomain = strings.Index(email, "@") + 1
	poinTLD = strings.Index(email, ".") + 1
	last = strings.LastIndex(email, "")

	for poinDomain < poinTLD-1 {
		domain += string(email[poinDomain])
		poinDomain++
	}
	for poinTLD < last {
		TLD += string(email[poinTLD])
		poinTLD++
	}

	result := "Domain: " + domain + " dan TLD: " + TLD
	return result // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.com"))
}
