package main

import "fmt"

func main() {

	friends := []string{"Marry", "John", "Paul", "Diana"}
	buddies := []string{}
	buddies = append(buddies, friends...)

	buddies[3] = "Dan"

	fmt.Println("Friends:", friends)
	fmt.Println("Buddies:", buddies)
}
