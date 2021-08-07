package main

import "fmt"

const last_name string = "Sim"

func plus(a int, b int, name string) (int, string) {
	return a + b, name
}

func plusInRange(a ...int) int {
	total := 0
	for _, item := range a {
		total += item
	}
	return total
}

func printName() {
	name := "Mason !!!!!!! Is my name"

	for index, item := range name {
		fmt.Println(index, string(item))
	}
}

func main() {
	fmt.Println("Welcome to Coin world")
	result, ret_name := plus(2, 2, last_name)
	fmt.Println(result, ret_name)

	result_number := plusInRange(1, 3, 4, 2, 5, 23, 3, 4, 5, 54)
	fmt.Println(result_number)

	printName()
}
