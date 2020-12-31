package main

import "fmt"

func main() {
	a := 10
	b := 1.5

	a = int(b) // explicit conversion or casting is needed. otherwise, it is a compiler error
	fmt.Println(a, b)

	var count int   // 0
	var price float64  // 0
	var name string  // 'space'
	var onSale bool // false

	fmt.Println(count, price, name, onSale)
}
