package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	totalVIP := VIP * 30
	totalRegular := regular * 20
	totalStudent := student * 10

	price := float32(totalVIP + totalRegular + totalStudent)

	if price >= 100 {
		if day%2 != 0 {
			if VIP+regular+student < 5 {
				price = price - (price * 0.15)
			} else {
				price = price - (price * 0.25)
			}
		} else {
			if VIP+regular+student < 5 {
				price = price - (price * 0.1)
			} else {
				price = price - (price * 0.2)
			}
		}
	}

	return price // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(4, 4, 4, 22))
}
