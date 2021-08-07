package main

import "fmt"

func main() {
	foods := []string{"potato", "pizza", "pasta"}
	fmt.Printf("%v\n", foods)
	// for _, food := range foods {
	// 	fmt.Println(food)
	// }
	foods = append(foods, "tomato")
	fmt.Printf("%v\n", foods)
}
