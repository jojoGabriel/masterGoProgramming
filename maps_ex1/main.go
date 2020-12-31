package main

import "fmt"

func main() {

	var m1 map[string]string

	fmt.Printf("m1 %#v\n", m1)

	m2 := map[int]string{1: "one", 2: "two"}
	m2[10] = "Abba"

	fmt.Printf("Existing: %v\n", m2[2])
	fmt.Printf("Not Existing: %v\n", m2[100])

}
