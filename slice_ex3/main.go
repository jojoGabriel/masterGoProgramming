package main

import "fmt"

func main() {
	nums := []float64{1.25, 2.50, 3.75}

	nums = append(nums, 10.1)

	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println(nums)

	n := []float64{2.3, 7.2}

	nums = append(nums, n...)

	fmt.Println(nums)

}
