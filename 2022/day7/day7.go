package main

import (
	"fmt"
)

func main() {
	// PART 1
 	test := parseFile("input.txt")
   
	fmt.Printf("(Part 1) Start of packet after: %d chars \n \n", test)

	// PART 2
	//fmt.Printf("(Part 2) Start of packet after: %d chars \n \n", )
}

func parseFile(filename string) int {
        var test int
	if filename == "input.txt" {
           test = 4     
	} else if filename == "dumb" {
           test = 5
 	} else if filename == "input.txt" {
	   test = 9
	}

       return test
}
