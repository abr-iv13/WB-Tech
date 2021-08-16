package main

import (
	"log"

	wallet "github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/facade/wallet"
)

func main() {
	//Инициализация экземплярa класса.
	walletFacade := wallet.NewWalletFacade("abc", 1234)

	//Добавление средств в кошелек
	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	//Списание средств с кошелька.
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
