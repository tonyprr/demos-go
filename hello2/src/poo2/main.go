package main

import (
	"fmt"
	fig "poo2/figuras"
)

func main() {

	myCuadrado := fig.Cuadrado{Base: 2}
	myRectangulo := fig.Rectangulo{Base: 2, Altura: 3}

	fig.Calcular(myCuadrado)
	fig.Calcular(myRectangulo)

	// Lista interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
