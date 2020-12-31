package main

import "fmt"

func main() {
	nums := []int{30, -1, -6, 90, -6}

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i]%2 == 0 {
			count++
		}
	}
	fmt.Println("positive even: ", count)
}
