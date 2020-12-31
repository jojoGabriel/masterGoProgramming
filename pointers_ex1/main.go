package main

import "fmt"

func main() {

	x := 10.10
	fmt.Println("address of x:", &x)

	ptr := &x
	fmt.Printf("ptr type: %T, value: %v\n", ptr, ptr)

	fmt.Printf("ptr address %p,\nvalue of x through ptr %v\n", &ptr, *ptr)

}
