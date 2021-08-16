package main

import (
	"fmt"
	"log"
)

func main() {
	vendingMachine := vending.NewVendingMachine(1, 10)

	err := vending.VendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vending.VendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vending.VendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vending.VendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vending.VendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vending.VendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vending.VendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
