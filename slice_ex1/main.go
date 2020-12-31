package main

import "fmt"

func main() {

	names := []string{"Jon", "Joy", "Jah"}

	for i, name := range names {
		fmt.Printf("names[%d]=%s\n", i, name)
	}
}
