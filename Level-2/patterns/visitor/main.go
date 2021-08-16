package main

import (
	"fmt"

	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/visitor/figure"
)

func main() {
	square := &figure.Square{Side: 2}
	circle := &figure.Circle{Radius: 3}
	rectangle := &figure.Rectangle{L: 2, B: 3}

	areaCalculator := &figure.AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &figure.MiddleCoordinates{}
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
