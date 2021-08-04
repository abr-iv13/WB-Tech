//31: Написать программу нахождения расстояния между 2 точками,
//которые представление в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package main

import "fmt"

type Point struct {
	x, y int
}

func newPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p Point) showDistance() {
	fmt.Println("Расстояние:", p.x-p.y)
}

func main() {
	p := newPoint(27, 20)
	p.showDistance()
}
