package main

import (
	// "errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	// "strings"
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

	// Line Format is #x#x#
	fmt.Printf("%T\n", fileContents) // []uint8
	// Get line from file
	dimensions := ""
	var dimList []string
	for i, x := range fileContents {
		fmt.Printf("index[%d]\t%v\t%d\n", i, x, int(x))
		if x != 10 {
			dimensions += string(x)
			fmt.Printf("%s\n", dimensions)
		} else {
			fmt.Printf("%s\n", dimensions)
			dimList = append(dimList, dimensions)
			dimensions = ""
		}
	}
	fmt.Println(dimList)
	// Line -> String "#x#x#"
	// Split string by 'x'
	// Check if all # are numbers
	// Put into an int slice
	// Sort from least to greatest
	// Calculate surface area
	// Calculate slack
	// Sum

}
