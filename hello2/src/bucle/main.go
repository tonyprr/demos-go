package main

import "fmt"

func main() {

	// for condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Printf("\n")

	// for reverse
	for j := 20; j >= 0; j-- {
		if j%2 == 0 {
			fmt.Println(j)
		}
	}

	fmt.Printf("\n")

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condición
	value := 50
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condición")
	}

	// for forever
	/* 	counterForever := 0
	   	for {
	   		fmt.Println(counterForever)
	   		counterForever++
	   	}
	*/
}
