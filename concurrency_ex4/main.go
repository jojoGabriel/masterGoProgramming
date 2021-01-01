package main

import (
	"fmt"
)

func power(i int, c chan int) {
	c <- i * i
}

func main() {

	c := make(chan int)

	for i := 1; i <= 50; i++ {
		go power(i, c)
	}

	// time.Sleep(time.Second * 2)

	for i := 1; i <= 50; i++ {
		fmt.Printf("val = %v\n", <-c)
	}
}
