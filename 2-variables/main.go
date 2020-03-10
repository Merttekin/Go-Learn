package main

import "fmt"

func main() {
	var x int
	x = 10
	fmt.Printf("x = %d\n", x)

	a := 20
	fmt.Printf("a = %d\n", a)
	a = 30
	fmt.Printf("a = %d\n", a)

	name := "wilson"
	fmt.Println("Sa", name)

	pi := 3.14
	fmt.Printf("PI = %0.3f\n", pi)

	ilk := true
	son := false

	fmt.Println("Değerler eşit mi?(true/false) ", ilk == son)
}
