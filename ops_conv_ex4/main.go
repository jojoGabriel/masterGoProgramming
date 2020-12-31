package main

import "fmt"

func main() {

	sunToEarth := 149600000 * 1000
	speedOfLight := 299792458

	time := sunToEarth / speedOfLight

	fmt.Println(time)
}
