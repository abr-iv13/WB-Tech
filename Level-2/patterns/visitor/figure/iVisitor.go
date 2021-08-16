//Посетитель
package figure

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForrectangle(*Rectangle)
}
