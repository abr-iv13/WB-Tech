package businessLogic

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForrectangle(*Rectangle)
}
