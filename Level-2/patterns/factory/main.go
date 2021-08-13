package main

import (
	"fmt"

	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/facade/gun"
)

func main() {
	ak47, _ := gun.GetGun("ak47")
	musket, _ := gun.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g gun.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
