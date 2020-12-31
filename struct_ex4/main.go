package main

import "fmt"

func main() {

	type grades struct {
		grade  int
		course string
	}

	type person struct {
		name           string
		age            int
		favoriteColors []string
		grades         grades
	}

	me := person{
		name:           "Jo",
		age:            25,
		favoriteColors: []string{"blue", "black"},
		grades: grades{
			grade:  1,
			course: "golang",
		},
	}

	you := person{
		name:           "Joy",
		age:            25,
		favoriteColors: []string{"blue", "gray"},
		grades: grades{
			grade:  2,
			course: "go",
		},
	}

	fmt.Printf("%+v\n", me)

	fmt.Printf("%+v\n", you)

}
