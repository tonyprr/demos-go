package main

import (
	"fmt"

	math "github.com/tonyprr/libmath"
)

func main() {
	suma := math.Sum(2, 3, 40)
	multi := math.Multiply(2, 3, 40)

	fmt.Println(suma)
	fmt.Println(multi)
}
