package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// De esta manera aplicamos la composicion sobre la herencia
type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

// Creamos la interfaz
type PrintInfo interface {
	getMessage() string
}

// Composicion sobre herencia
func (ft FullTimeEmployee) getMessage() string {
	return fmt.Sprintf("Hi %s, you are %d years old. And you are a full time employee", ft.name, ft.age)
}

func (te TemporaryEmployee) getMessage() string {
	return fmt.Sprintf("Hi %s, you are %d years old. And you are a temprary employee", te.name, te.age)
}

// Creamos el metodo de la interfaz
func getMessage2(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Benjamin"
	ftEmployee.age = 20
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessage2(tEmployee)
	getMessage2(ftEmployee)
}
