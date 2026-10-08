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
	fmt.Printf("%T\n", fileContents) // []uint8
	// Get line from file
	// Line -> String "#x#x#"
	dimensions := ""
	var strDimList []string //= make([]string, 2, 4)
	for i, x := range fileContents {
		fmt.Printf("index[%d]\t%v\t%d\n", i, x, int(x))
		if x == 13 {
			continue
		} else if x != 10 {
			dimensions += string(x)
			fmt.Printf("Working: %s\n", dimensions)
		} else {
			fmt.Printf("Final: %s\n", dimensions)
			strDimList = append(strDimList, dimensions)
			dimensions = ""
		}
	}
	fmt.Println()
	for i, x := range strDimList {
		fmt.Printf("item[%d] in dimList: %v\n", i, x)
	}
	fmt.Println()

	// Split string by 'x'
	var strDimListList [][]string
	for i, item := range strDimList {
		lst := strings.Split(item, "x")
		strDimListList = append(strDimListList, lst)
		fmt.Printf("item[%d] in strDimListList: %v\n", i, lst)
	}
	fmt.Println(strDimListList)
	fmt.Println()

	// Check if all # are numbers
	// Put into an int slice
	var intDimListList [][]int
	for _, list := range strDimListList {
		var tempList []int
		for _, num := range list {
			// fmt.Printf("%v is a %T\n", num, num)
			x, err := strconv.Atoi(num)
			// fmt.Printf("%v is a %T\n", x, x)
			checkError(err)
			tempList = append(tempList, x)
		}
		intDimListList = append(intDimListList, tempList)
		// fmt.Printf("%v is a %T\n", num, num)

	}
	fmt.Println()
	fmt.Println()
	fmt.Println(intDimListList)
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
	fmt.Println(intDimListList)

	var runningSurfaceArea int
	for _, list := range intDimListList {
		// Calculate surface area
		listSurfaceArea := 2*list[0]*list[1] + 2*list[0]*list[2] + 2*list[1]*list[2]
		// Calculate slack
		listSlackArea := list[0] * list[1]
		fmt.Printf("Surface Area = %d\tSlack Area = %d\n", listSurfaceArea, listSlackArea)
		runningSurfaceArea += listSurfaceArea + listSlackArea
	}
	fmt.Println()
	fmt.Printf("sqft needed: %d", runningSurfaceArea)
	fmt.Println()

	a := 10
	b := 14
	fmt.Printf("a = %d\tb = %d\n", a, b)
	a, b = swapInt(a, b)
	fmt.Printf("a = %d\tb = %d\n", a, b)

	// Sum

}
