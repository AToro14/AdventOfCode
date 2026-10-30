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

func main() {
	fileIn := filepath.Join(filepath.Dir(os.Args[0]), "input.txt")
	fileContents, err := os.ReadFile(fileIn)
	checkError(err)
	currPosX := 0
	currPosY := 0
	rCurrPosX := 0
	rCurrPosY := 0
	visited := make(map[coords]struct{})
	housesVisited := 1
	currCoords := coords{x: currPosX, y: currPosY}
	// rCurrCords := coords{x: rCurrPosX, y: rCurrPosY}
	visited[currCoords] = struct{}{}
	instructionCount := 0
	for _, x := range fileContents {
		instructionCount++
		santa := instructionCount % 2
		if santa == 0 {
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
				instructionCount--
				continue
			}
			currCoords := coords{x: currPosX, y: currPosY}

			if _, ok := visited[currCoords]; !ok {
				visited[currCoords] = struct{}{}
				housesVisited++
			}
		} else {
			switch x {
			case 94:
				rCurrPosY++
			case 62:
				rCurrPosX++
			case 118:
				rCurrPosY--
			case 60:
				rCurrPosX--
			default:
				instructionCount--
				continue
			}
			rCurrCoords := coords{x: rCurrPosX, y: rCurrPosY}

			if _, ok := visited[rCurrCoords]; !ok {
				visited[rCurrCoords] = struct{}{}
				housesVisited++
			}
		}
	}
	fmt.Printf("Number of houses visited:\t%d\n", housesVisited)
}
