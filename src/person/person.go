package person

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) Set(name string, age int) {
	p.name = name;
	p.age = age;
}

func (p *Person) PrintInfo(){
	fmt.Printf("Hello!! my name is %s, and my age is %d\n", p.name, p.age);
}