package main

import "fmt"

func main() {
	var i int = 3
	var f float64 = 3.2

	j := float64(i)
	g := int(f)

	fmt.Printf("j = %f is a %T\n", j, j)
	fmt.Printf("g = %d is a %T\n", g, g)
}
