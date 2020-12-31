package main

import "fmt"

func main() {

	friends := []string{"Marry", "John", "Paul", "Diana"}
	buddies := make([]string, len(friends))
	copy(buddies, friends)

	friends[3] = "Dana"

	fmt.Println("Friends:", friends)
	fmt.Println("Buddies:", buddies)
}
