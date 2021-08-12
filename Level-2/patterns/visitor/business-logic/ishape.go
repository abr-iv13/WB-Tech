//Элемент
package businessLogic

type Shape interface {
	GetType() string
	Accept(Visitor)
}
