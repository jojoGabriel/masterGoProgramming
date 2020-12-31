package main

import "fmt"

func main() {

	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

	sum := 0
	for _, num := range nums[1 : len(nums)-2] {
		sum += num
	}
	fmt.Println(nums, sum)
}
