package main

import "fmt"

func main() {

	var c1 chan float64

	c2 := make(<-chan rune)

	c3 := make(chan<- rune)

	c4 := make(chan int, 10)

	fmt.Printf("c1 type %T value, %v\n", c1, c1)

	fmt.Printf("c2 type %T value, %v\n", c2, c2)

	fmt.Printf("c3 type %T value, %v\n", c3, c3)

	fmt.Printf("c4 type %T value, %v\n", c4, c4)
}
