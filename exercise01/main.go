package main

import "fmt"

var pkgvar int = 1

func main() {

	var locvar int = 2
	fmt.Println("pkgvar = ", pkgvar)
	fmt.Println("locvar = ", locvar)
}
