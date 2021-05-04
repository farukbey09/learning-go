package main

import "fmt"

type Person struct {
	name string
	age  int
}

func createPerson(name string, age int) *Person {
	p := Person{
		name: name,
		age:  age,
	}

	return &p
}
func main() {
	person := createPerson("Ahmet", 20)
	fmt.Println(person)
	fmt.Println(person.name, person.age)

}
