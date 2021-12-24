package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (somePC pc) ping() {
	fmt.Println(somePC.brand, "Pong")
}

func (somePC *pc) duplicateRAM() {
	somePC.ram = somePC.ram * 2
}

func (somePC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", somePC.ram, somePC.disk, somePC.brand)
}

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
}
