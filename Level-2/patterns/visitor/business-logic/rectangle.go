package businessLogic

type Rectangle struct {
	L int
	B int
}

func (t *Rectangle) Accept(v Visitor) {
	v.visitForrectangle(t)
}

func (t *Rectangle) GetType() string {
	return "rectangle"
}
