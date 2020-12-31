package main

import (
	"log"
	"os"
)

func main() {

	createFile("info.txt")

}

func createFile(fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
