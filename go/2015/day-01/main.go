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
	fileIn := filepath.Join(filepath.Dir(os.Args[0]), "test.txt")
	fileContents, err := os.ReadFile(fileIn)
	checkError(err)

	floorNum := 0
	basementPosition := 0

	for i, x := range fileContents {
		switch x {
		case 40: // '('
			floorNum++
		case 41: // ')'
			floorNum--
		default:
			checkError(nil)
		}
		if basementPosition == 0 && floorNum == -1 {
			basementPosition = i + 1
		}
	}

	fmt.Printf("Final Floor Number:\t\t%d\n", floorNum)
	fmt.Printf("First Basement Position:\t%d\n", basementPosition)
}
