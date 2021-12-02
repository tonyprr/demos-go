package main

import "fmt"

type Animal struct {
	Id   int
	Name string
}

// Constructor
func NewAnimal(id int, name string) *Animal {
	return &Animal{
		Id:   id,
		Name: name,
	}
}

func (e *Animal) SetName(name string) {
	e.Name = name
}

func (e *Animal) GetName() string {
	return fmt.Sprintf("this is my name: %s", e.Name)
}

func main() {
	animal := NewAnimal(1, "New Animal")
	fmt.Println(animal.Id)
	// 1
	fmt.Println(animal.Name)
	// NewAnimal
	fmt.Println("Rename Animal")
	animal.SetName("New Animal Name")
	fmt.Println(animal.Name)
	// New Animal Name
	fmt.Println(animal.GetName())
	// this is my name: New Animal Name
}
