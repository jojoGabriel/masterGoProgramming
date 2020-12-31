package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func main() {

	var price money = 10.50
	price.print()

	var salary money = 5000.
	money.print(salary)
}
