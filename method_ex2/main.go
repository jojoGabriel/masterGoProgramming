package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func main() {

	var price money = 10.50
	price.print()

	fmt.Printf("PriceStr %s", price.printStr())

}
