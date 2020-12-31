package main

import (
	"fmt"
	"strings"
)

func searchItem(list []string, item string) bool {

	found := false
	for _, v := range list {
		if strings.EqualFold(v, item) {
			found = true
			break
		}
	}
	return found
}

func main() {
	animals := []string{"Lion", "tiger", "bear"}
	result := searchItem(animals, "beaR")
	fmt.Println(result) // => true
	result = searchItem(animals, "lion")
	fmt.Println(result) // => false
}
