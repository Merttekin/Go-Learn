package main

import "fmt"

func main() {
	people := map[string]int{
		"wilson": 25,
		"andre":  26,
	}

	people["junior"] = 10

	for name, age := range people {
		fmt.Printf("Name: %s yaş: %d\n", name, age)
	}

	delete(people, "wilson")
	fmt.Printf("Wlson: %d\n", people["wilson"])
}
