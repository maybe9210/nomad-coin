package main

import "fmt"

type person struct {
	name string
	age int
}

func (p person) sayHello() {
	fmt.Printf("Hello! my name is %s and I'm %d\n", p.name, p.age)
}

func main() {
	hulo := person{age: 12, name: "hulo"}
	hulo2 := person{"hulo2", 13}
	hulo.sayHello()
	hulo2.sayHello()
}
