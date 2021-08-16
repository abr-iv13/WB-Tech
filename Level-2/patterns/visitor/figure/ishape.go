//Элемент
package figure

type Shape interface {
	GetType() string
	Accept(Visitor)
}
