package person

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) SetDetail(name string, age int) {
	p.name = name
	p.age = age
	fmt.Printf("hello my name is %s, and my age is %d\n", name, age)
}
