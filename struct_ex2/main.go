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

	me.name = "Andrei"

	you := person{
		name:           "Joy",
		age:            25,
		favoriteColors: []string{"blue", "gray"},
	}

	colors := you.favoriteColors
	you.favoriteColors = append(you.favoriteColors, "green")

	fmt.Printf("colors is %T = %v\n", colors, colors)

	fmt.Printf("%+v\n", me)

	fmt.Printf("%+v\n", you)

}
