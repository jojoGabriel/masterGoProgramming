package main

import "fmt"

func main() {

	x, y := 10, 2
	ptrX, ptrY := &x, &y

	z := *ptrX / *ptrY

	fmt.Println(z)
}
