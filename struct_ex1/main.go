package main

import "fmt"

func main() {

	type person struct {
		name           string
		age            int
		favoriteColors []string
	}

	me := person{
		name:           "Jo",
		age:            25,
		favoriteColors: []string{"blue", "black"},
	}

	you := person{
		name:           "Joy",
		age:            25,
		favoriteColors: []string{"blue", "gray"},
	}

	fmt.Printf("%+v\n", me)

	fmt.Printf("%+v\n", you)

}
