package main

import "fmt"

func sum(n ...int) (sum int) {

	sum = 0
	for _, v := range n {
		sum += v
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(2, 2, 2, 2, 2))

}
