package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("x = %d, y = %f, z = %s\n", x, y, z)
	fmt.Printf("score = %#v\n", score)

	fmt.Printf("z = %q\n", z)

	fmt.Printf("x = %v, y = %v, z = %v\n", x, y, z)
	fmt.Printf("score = %v\n", score)

	fmt.Printf("y is a %T and z is a %T", y, score)
}
