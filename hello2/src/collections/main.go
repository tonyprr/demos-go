package main

import (
	"fmt"
	"strings"
)

func main() {
	// basicArrayAndSlice()
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20
	m["Juan"] = 2
	m["Ana"] = 30
	m["Antonio"] = 40

	fmt.Println(m)

	//Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}

func basicArrayAndSlice() {
	// Array
	var array [4]int
	array[0] = 1
	array[3] = 10
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	// basicArrayAndSlice()
	slice1 := []string{"hola", "que", "hace"}
	for i, value := range slice1 {
		fmt.Println(i, value)
	}

	//ama
	//amor a roma
	text := "Ama"
	isPalindromo(text)
	fmt.Println(text)

	slice2 := []string{"hola", "que", "hace"}
	for i, value := range slice2 {
		fmt.Println(i, value)
	}

	//ama
	//amor a roma
	text2 := "Ama"
	isPalindromo(text2)
	fmt.Println(text2)

}

func isPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}
