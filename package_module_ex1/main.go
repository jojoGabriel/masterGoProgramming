package main

import (
	"fmt"
	"packages/mypackages/arithmetic"
)

func main() {
	fmt.Println("13 is prime =", arithmetic.IsPrime(13))
	fmt.Println("25 is prime =", arithmetic.IsPrime(25))
	fmt.Println("5 factorial = ", arithmetic.Factorial(5))
}
