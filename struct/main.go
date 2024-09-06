package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 10

	return &p
}

func main() {
	fmt.Println(Person{"funminiyi", 8})

	fmt.Println(Person{name: "rere", age: 2})

	fmt.Println(Person{name: "kudos"})

	fmt.Println(&Person{name: "ann", age: 40})

	fmt.Println(newPerson("young"))

	age := newPerson("young").age
	age = 11
	fmt.Println(age)

	s := Person{name: "sean", age: 50}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"t-rex baby",
		true,
	}
	fmt.Println(dog)
}
