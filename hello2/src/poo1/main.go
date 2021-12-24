package main

import (
	pk "demos-go/hello2/src/poo1/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	e := pk.Employee{}
	fmt.Println(e)
	e.SetId(5)
	e.SetName("Luis")
	fmt.Println(e.GetId(), e.GetName())

	e2 := new(pk.Employee)
	fmt.Printf("%v\n", &e2)
	fmt.Printf("%v\n", *e2)
	e2.SetId(10)
	e2.SetName("Mario")
	fmt.Printf("%v\n", *e2)

	e3 := pk.NewEmployee(1, "Antonio", true)
	fmt.Println(*e3)

}
