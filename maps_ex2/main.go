package main

import "fmt"

func main() {
	m1 := map[int]bool{}
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	s2 := fmt.Sprintf("%s", m2)
	s3 := fmt.Sprintf("%s", m3)

	fmt.Println(s2, s3)
	fmt.Println(s2 == s3)
}
