package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please pass 2-10 numbers")
		return
	}

	sum := 0.
	product := 1.
	for _, v := range os.Args[1:] {
		nbr, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}
		sum += nbr
		product *= nbr
	}
	fmt.Printf("Sum: %v, Product: %v\n", sum, product)

}
