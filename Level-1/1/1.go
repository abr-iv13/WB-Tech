// 1: Реализовать композицию структуры Action от родительской структуры Human.
package main

import "fmt"

type (
	Human struct {
		lastName  string
		firstName string
		surName   string
	}

	Action struct {
		Human
	}
)

func (h Human) FullName() Human {
	return h
}

func (h *Human) SetName(l string, f string, s string) {
	h.lastName = l
	h.firstName = f
	h.surName = s
}

func main() {
	// action := &Action{} - Синтаксический сахар для new(Action)
	action := new(Action)
	action.SetName("Ivanov", "Ivan", "Ivanovich")
	fmt.Println(action.FullName())
}
