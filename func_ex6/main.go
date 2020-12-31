package main

import "fmt"

func searchItem(list []string, item string) bool {

	found := false
	for _, v := range list {
		if v == item {
			found = true
			break
		}
	}
	return found
}

func main() {
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result) // => false
}
