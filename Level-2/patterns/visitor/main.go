package main

import (
	"fmt"

	bussinesLogic "github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/visitor/business-logic"
)

func main() {
	square := &bussinesLogic.Square{Side: 2}
	circle := &bussinesLogic.Circle{Radius: 3}
	rectangle := &bussinesLogic.Rectangle{L: 2, B: 3}

	areaCalculator := &bussinesLogic.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &bussinesLogic.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
