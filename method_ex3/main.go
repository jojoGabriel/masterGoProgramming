package main

import "fmt"

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price = b.price * 0.9
}

func main() {
	bk := book{price: 9.9, title: "golang"}

	fmt.Printf("vat: %.2f\n", bk.vat())

	bk.discount()

	fmt.Println(bk)
}
