package main

import "os"

func main() {

	if filePresent("info.txt") {
		_ = renameFile("info.txt", "data.txt")
	}

}

func filePresent(filepath string) bool {

	_, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func renameFile(oldName string, newName string) bool {

	err := os.Rename(oldName, newName)
	if err != nil {
		return false
	}

	return true
}
