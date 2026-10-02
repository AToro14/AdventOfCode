package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)
	fmt.Println(os.Args[0]) // This file
	fileIn := filepath.Join(filepath.Dir(os.Args[0]), "test.txt")
	fmt.Println(fileIn)
	fileContents, err := os.ReadFile(fileIn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(fileContents))
}
