package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Printf("value1 = %d, value2 = %d\n", value1, value2)

	value1, _ = doubleReturn(5)
	fmt.Printf("value1 = %d", value1)

}
