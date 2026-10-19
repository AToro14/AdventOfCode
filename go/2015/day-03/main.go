// day 03 solution
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	// "strconv"
	// "strings"
)

// map of coordinates to empty struct
type coords struct {
	x int
	y int
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func parseAngleBrackets() {}

func main() {
	fileIn := filepath.Join(filepath.Dir(os.Args[0]), "input.txt")
	fileContents, err := os.ReadFile(fileIn)
	checkError(err)
	currPosX := 0
	currPosY := 0
	visited := make(map[coords]struct{})
	housesVisited := 1
	currCoords := coords{x: currPosX, y: currPosY}
	visited[currCoords] = struct{}{}
	for _, x := range fileContents {
		fmt.Printf("%d\n", x)
		// replace with switch case
		switch x {
		case 94:
			currPosY++
		case 62:
			currPosX++
		case 118:
			currPosY--
		case 60:
			currPosX--
		default:
			continue
		}
		currCoords := coords{x: currPosX, y: currPosY}
		fmt.Println(currCoords)
		fmt.Println(currPosX, "\t", currPosY)

		if _, ok := visited[currCoords]; !ok {
			visited[currCoords] = struct{}{}
			housesVisited++
		}
		// fmt.Println(visited)
	}
	fmt.Printf("Number of houses visited:\t%d\n", housesVisited)
	// fmt.Println(visited)
}

// direction inc or dec the current position
// attempt to add current position to map
// if the same don't add, we've been there before
