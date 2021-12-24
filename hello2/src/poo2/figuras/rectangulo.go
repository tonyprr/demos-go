package figuras

type Rectangulo struct {
	Base   float64
	Altura float64
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}
