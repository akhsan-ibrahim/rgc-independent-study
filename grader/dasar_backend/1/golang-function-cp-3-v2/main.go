package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	var letter string
	for i := 65; i < 90; i++ {
		letter += string(byte(i))
	}
	var sep string
	for i := 0; i < len(names); i++ {
		if strings.Contains(strings.ToLower(letter), strings.ToLower(string(names[i]))) == false {
			sep = string(names[i])
		}
	}

	namesWord := strings.Split(names, sep)
	shortest := 100001
	var lsShortestName []string
	for i := 0; i < len(namesWord); i++ {
		if len(namesWord[i]) < shortest {
			shortest = len(namesWord[i])
		}
	}
	for i := 0; i < len(namesWord); i++ {
		if len(namesWord[i]) == shortest {
			lsShortestName = append(lsShortestName, string(namesWord[i]))

		}
	}
	var shortestName string
	if len(lsShortestName) == 1 {
		shortestName = string(lsShortestName[0])
	} else {
		var lsn []int
		min := 100001
		for i := 0; i < len(lsShortestName); i++ {
			var n int
			for j := 0; j < len(lsShortestName[i]); j++ {
				n += int(lsShortestName[i][j])
			}
			lsn = append(lsn, n)
		}
		for i := 0; i < len(lsn); i++ {
			if lsn[i] < min {
				min = lsn[i]
				shortestName = string(lsShortestName[i])
			}
		}
	}

	return shortestName // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Hanif Joko"))
	fmt.Println(FindShortestName("Hanif;Joko;Indah;Ari;Intan"))
	fmt.Println(FindShortestName("Ari,Aru,Ara,Andi,Asik"))
	fmt.Println(FindShortestName("A,B,C,D,E"))
	fmt.Println(FindShortestName("Budi;Tia;Tio")) // "Tia"

}
