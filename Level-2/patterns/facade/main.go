package main

import (
	"fmt"
	"log"

	patternFacade "github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/facade/pattern-facade"
)

func main() {
	fmt.Println()
	walletFacade := patternFacade.NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
