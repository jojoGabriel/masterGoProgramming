package main

import (
	"fmt"
)

func main() {

	r := 'ă'

	fmt.Printf("r type of %T\n", r)

	ma := "ma"
	m := "m"

	mama := ma + m + string(r)

	fmt.Println(mama)
}
