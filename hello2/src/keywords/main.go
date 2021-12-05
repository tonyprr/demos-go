package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {

		// Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		fmt.Println(i)

		if i == 8 {
			fmt.Println("break")
			break
		}
	}

}
