package main

import "fmt"

func main() {

	var empty interface{}

	empty = 1
	fmt.Printf("%T\n", empty)

	empty = 10.
	fmt.Printf("%T\n", empty)

	empty = []int{1, 10}
	fmt.Printf("%T\n", empty)

	empty = append(empty.([]int), 20)
	fmt.Println(empty)

}
