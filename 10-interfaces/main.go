package main

import "fmt"

type Animal interface {
	Shout() string
}

type Cat struct{}

func (c *Cat) Shout() string {
	return "Cat"
}

type Dog struct{}

func (d *Dog) Shout() string {
	return "Dog"
}

func main() {
	var animal Animal

	animal = &Cat{}
	callMyAnimal(animal)

	animal = &Dog{}
	callMyAnimal(animal)
}

func callMyAnimal(a Animal) {
	fmt.Println("Animal:", a.Shout())
}
