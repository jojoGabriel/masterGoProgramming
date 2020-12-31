package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	is := strconv.Itoa(i)
	fmt.Printf("is type %T value %s\n", is, is)

	s2i, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("s2i type %T value %d\n", s2i, s2i)
	} else {
		fmt.Println("Unable to convert s2 to int")
	}

	fs := strconv.FormatFloat(f, 'f', 1, 64)
	fmt.Printf("fs type %T value %q\n", fs, fs)

	s1f, err := strconv.ParseFloat(s1, 64)
	if err == nil {
		fmt.Printf("s1f type %T value %.2f", s1f, s1f)
	} else {
		fmt.Println("Unable to convert s1 to float")
	}

}
