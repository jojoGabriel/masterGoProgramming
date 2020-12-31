package main

import "fmt"

func main() {

	type duration int
	var hour duration

	fmt.Printf("hour = %v\n", hour)
	fmt.Printf("hour is a %T\n", hour)

	hour = 3600

	fmt.Printf("hour = %v\n", hour)
}
