package main

import "fmt"

func main() {

	var cities [4]string
	fmt.Printf("cities %#v\n", cities)

	grades := [3]float64{1.2, 2.3}
	fmt.Printf("grades %#v\n", grades)

	taskDone := [...]bool{}
	fmt.Printf("taskDone %v\n", taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Printf("cities[%d] = %s\n", i, cities[i])
	}

	for i, grade := range grades {
		fmt.Printf("grades[%d] = %f\n", i, grade)
	}
}
