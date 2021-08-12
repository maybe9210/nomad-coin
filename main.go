package main

import "fmt"

func main() {
	a:=2
	b:=&a
	a=5
	fmt.Println(*b)
}
