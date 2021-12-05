package main

import "fmt"

func main() {
	// Declaración de variables
	helloMessage := "Hello"
	worldMesage := "world"

	// Println
	fmt.Println(helloMessage, worldMesage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
