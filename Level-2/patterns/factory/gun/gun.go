package gun

type Gun struct {
	Name  string
	Power int
}

func (g *Gun) SetName(name string) {
	g.Name = name
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) setPower(power int) {
	g.Power = Power
}

func (g *Gun) getPower() int {
	return g.Power
}
