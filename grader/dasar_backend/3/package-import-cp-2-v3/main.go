package main

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/service"
	"fmt"
)

func CashierApp(db *database.Database) service.ServiceInterface {
	service := service.NewService(db)

	return service
}

// gunakan untuk debugging
func main() {
	db := database.NewDatabase()
	serv := CashierApp(db)

	serv.AddCart("Kaos Polos", 2)
	fmt.Println(serv.ShowCart())
	serv.AddCart("Kaos sablon", 1)
	fmt.Println(serv.ShowCart())

	err := serv.AddCart("Topi", 3)
	if err != nil {
		// panic(err)
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
	fmt.Println(serv.ShowCart())

	err2 := serv.AddCart("Kaos Polos", 0)
	if err2 != nil {
		// panic(err2)
		fmt.Println(err2)
	} else {
		fmt.Println("success")
	}
	fmt.Println(serv.ShowCart())

	errPolos := serv.AddCart("Kaos sablon", -3)
	if errPolos != nil {
		// panic(err2)
		fmt.Println(errPolos)
	} else {
		fmt.Println("success")
	}

	fmt.Println(serv.ShowCart())
	// err3 := serv.RemoveCart("Kaos sablon")
	// if err3 != nil {
	// 	fmt.Println(err3)
	// } else {
	// 	fmt.Println("1 item removed")
	// }
	// fmt.Println(serv.ShowCart())

	// err4 := serv.ResetCart()
	// if err4 != nil {
	// 	fmt.Println(err4)
	// } else {
	// 	fmt.Println("reseted")
	// }
	// fmt.Println(serv.ShowCart())

	// prod, _ := serv.GetAllProduct()
	// fmt.Println(serv.GetAllProduct())
	// fmt.Println(len(prod))

	// fmt.Println(serv.Pay(50000))
	// fmt.Println(serv.ShowCart())
}
