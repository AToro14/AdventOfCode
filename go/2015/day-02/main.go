package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	fileIn := filepath.Join(filepath.Dir(os.Args[0]), "input.txt")
	fileContents, err := os.ReadFile(fileIn)
	checkError(err)

	// Line Format is #x#x#
	// Get line from file
	// Line -> String "#x#x#"
	// Split string by 'x'
	// Check if all # are numbers
	// Put into an int slice
	// Sort from least to greatest
	// Calculate surface area
	// Calculate slack
	// Sum

}
