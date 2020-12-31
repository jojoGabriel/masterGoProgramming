package main

import "fmt"

func cube(f float64) float64 {
	return f * f * f
}
func main() {
	fmt.Println(cube(3))
}
