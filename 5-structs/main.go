package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	p := Person{
		Name:   "Lorem Ipsum",
		Age:    26,
		Active: true,
	}
	fmt.Println("Name", p.Name)
	fmt.Println("Age", p.Age)
	fmt.Println("Active", p.Active)
}
