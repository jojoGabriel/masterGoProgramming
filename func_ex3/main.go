package main

import (
	"fmt"
	"strconv"
)

func myFunc(i string) int {

	n, err := strconv.Atoi(i)
	if err != nil {
		fmt.Println("Invalid number", i)
		return 0
	}

	nn, err := strconv.Atoi(i + i)
	nnn, err := strconv.Atoi(i + i + i)

	_ = err

	return (n + nn + nnn)
}

func main() {
	fmt.Println(myFunc("5"))
	fmt.Println(myFunc("9"))
	fmt.Println(myFunc("a"))
}
