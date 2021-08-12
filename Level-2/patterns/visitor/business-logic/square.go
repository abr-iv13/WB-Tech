package businessLogic

type Square struct {
	Side int
}

func (s *Square) accept(v Visitor) {
	v.VisitForSquare(s)
}

func (s *Square) GetType() string {
	return "Square"
}
