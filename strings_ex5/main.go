package main

import "fmt"

func main() {

	s := "你好 Go!"

	b := []byte(s)

	fmt.Printf("%#v\n", b)

	for i, v := range s {
		fmt.Printf("index: %d value: %d %c\n", i, v, v)
	}
}
