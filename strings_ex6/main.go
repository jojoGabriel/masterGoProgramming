package main

import "fmt"

func main() {

	s := "你好 Go!"

	r := []rune(s)

	fmt.Printf("rune %#v\n", r)

	for i, v := range r {
		fmt.Printf("index %d, value %d char %c\n", i, v, v)
	}
}
