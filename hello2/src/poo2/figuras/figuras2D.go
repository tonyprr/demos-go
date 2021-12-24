package figuras

import "fmt"

type Figuras2D interface {
	Area() float64
}

func Calcular(f Figuras2D) {
	fmt.Println("Area:", f.Area())
}
