package main

import "fmt"

func main() {

	for i := 1; i < 50; i++ {
		if i%7 == 00 {
			fmt.Println(i)
		}
	}
}
