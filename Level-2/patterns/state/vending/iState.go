//Интерфейс состояния
package vending

type State interface {
	AddItem(int) error
	RequestItem() error
	InsertMoney(money int) error
	DispenseItem() error
}
