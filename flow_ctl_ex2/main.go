package main

import "fmt"

func main() {

	i := 0
	for i < 50 {
		i++
		if i%7 != 00 {
			continue
		}
		fmt.Println(i)
	}
}
