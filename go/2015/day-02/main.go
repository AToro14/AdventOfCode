package main

import (
	// "errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func swapInt(a, b int) (A, B int) {
	a = a ^ b
	b = b ^ a
	a = a ^ b
	return a, b
}

func main() {
	fileIn := filepath.Join(filepath.Dir(os.Args[0]), "input.txt")
	fileContents, err := os.ReadFile(fileIn)
	checkError(err)

	// Line Format is #x#x#
	dimensions := ""
	var strDimList []string
	for _, x := range fileContents {
		if x == 13 { // cr carriage return
			continue
		} else if x != 10 { // lf line feed
			dimensions += string(x)
		} else {
			strDimList = append(strDimList, dimensions)
			dimensions = ""
		}
	}

	var strDimListList [][]string
	for _, item := range strDimList {
		lst := strings.Split(item, "x")
		strDimListList = append(strDimListList, lst)
	}

	// Check if all # are numbers
	// Put into an int slice
	var intDimListList [][]int
	for _, list := range strDimListList {
		var tempList []int
		for _, num := range list {
			x, err := strconv.Atoi(num)
			checkError(err)
			tempList = append(tempList, x)
		}
		intDimListList = append(intDimListList, tempList)

	}

	// Sort from least to greatest
	for _, list := range intDimListList {
		if list[0] > list[1] {
			list[0], list[1] = swapInt(list[0], list[1])
		}
		if list[0] > list[2] {
			list[0], list[2] = swapInt(list[0], list[2])
		}
		if list[1] > list[2] {
			list[1], list[2] = swapInt(list[1], list[2])
		}
	}

	var runningSurfaceArea int
	var runningRibbonLength int
	for _, list := range intDimListList {
		// Calculate surface area
		listSurfaceArea := 2*list[0]*list[1] + 2*list[0]*list[2] + 2*list[1]*list[2]
		// Calculate slack
		listSlackArea := list[0] * list[1]
		// Calculate ribbon length
		listRibbonLen := 2*list[0] + 2*list[1]
		// Calculate bow length
		listBowLen := list[0] * list[1] * list[2]
		runningSurfaceArea += listSurfaceArea + listSlackArea
		runningRibbonLength += listRibbonLen + listBowLen
	}

	fmt.Println()
	fmt.Printf("sqft needed: %d\n", runningSurfaceArea)
	fmt.Printf("len needed: %d\n", runningRibbonLength)
	fmt.Println()
}
