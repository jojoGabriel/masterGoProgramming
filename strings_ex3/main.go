package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "țară means country in Romanian"

	fmt.Println("Byte by byte")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v ", s1[i])
	}

	fmt.Println("\n" + strings.Repeat("*", 10))

	fmt.Println("Rune by rune")
	for _, char := range s1 {
		fmt.Printf("%c ", char)
	}

	fmt.Println("\n" + strings.Repeat("*", 10))

}
