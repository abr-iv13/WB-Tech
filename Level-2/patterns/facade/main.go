package main

import (
	"log"

	patternFacade "github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/facade/pattern-facade"
)

func main() {
	//Инициализация экземплярa класса.
	walletFacade := patternFacade.NewWalletFacade("abc", 1234)

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
