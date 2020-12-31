package main

import "fmt"

func f1(n uint) (int, int) {

	f := 1
	s := 0

	for i := 1; i <= int(n); i++ {
		f *= i
		s += i
	}

	return f, s
}

func main() {
	fmt.Println(f1(2))
}
