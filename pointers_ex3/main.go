package main

import "fmt"

func swap(i *float64, j *float64) {
	*i, *j = *j, *i
}

func main() {
	x, y := 5.5, 8.8
	swap(&x, &y)
	fmt.Println(x, y)
}
