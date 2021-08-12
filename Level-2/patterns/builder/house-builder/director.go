package houseBuilder

type Director struct {
	Builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		Builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.Builder = b
}

func (d *Director) BuildHouse() House {
	d.Builder.SetDoorType()
	d.Builder.SetWindowType()
	d.Builder.SetNumFloor()
	return d.Builder.GetHouse()
}
