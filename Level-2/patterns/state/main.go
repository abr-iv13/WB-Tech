package main

import (
	"fmt"
	"log"

	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/state/vending"
)

func main() {
	vendingMachine := vending.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
