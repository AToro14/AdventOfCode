package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	fmt.Println(os.Args[0]) // This file's path
	fileIn := filepath.Join(filepath.Dir(os.Args[0]), "test.txt")
	fmt.Println(fileIn) // Test file's path
	fileContents, err := os.ReadFile(fileIn)
	checkError(err)

	floorNum := 0
	for _, x := range fileContents {
		switch x {
		case 40: // '('
			floorNum++
		case 41: // ')'
			floorNum--
		default:
			checkError(nil)
		}
	}
	fmt.Printf("Floor Number: %d\n", floorNum)
}
