package main

import "fmt"

func main() {

	channel := make(chan string)

	go func(s string) {
		channel <- s
	}("golang")

	fmt.Printf("message: %s", <-channel)
}
