package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	if p.Age < 18 {
		return fmt.Sprintf("18 yaşından küçük olanlar: %s", p.Name)
	} else if p.Age > 50 {
		return fmt.Sprintf("50 yaşından büyük olanlar: %s", p.Name)
	}

	return fmt.Sprintf("Diğerleri: %s", p.Name)
}

func (p *Person) privateMethod() int {
	return p.Age * 4
}

func main() {
	person := Person{Name: "junior", Age: 17}
	fmt.Println(person.String())

	person = Person{Name: "Wilson", Age: 27}
	fmt.Println(person.String())

	person = Person{Name: "Antonio", Age: 60}
	fmt.Println(person.String())
}
