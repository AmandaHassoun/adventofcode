package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var slopes [][]string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dir, err := os.Getwd()
	check(err)
	file, err := os.Open(dir + "/day-3/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		slope := strings.Split(scanner.Text(), "")
		slopes = append(slopes, slope)
	}
	part1(slopes)

	//Part 2
	slope1 := multiply(1,1, slopes)
	slope2 := multiply(3,1, slopes)
	slope3 :=multiply(5,1, slopes)
	slope4 := multiply(7,1, slopes)
	slope5 :=multiply(1,2, slopes)

	fmt.Printf("(PART 2) Total multiplied answer: %d\n", slope1*slope2*slope3*slope4*slope5)
}

func part1(slopes [][]string){
	var trees = 0
	x := 0
	y := 0
	down := 1
	right := 3
	for y < len(slopes) {
		if slopes[y][x] == "#" {
			trees += 1
		}
		y += down
		x = (x + right)%len(slopes[0])
	}
	fmt.Printf("(PART 1) Total count of trees: %d\n\n", trees)
}

func multiply(right int, down int, slopes [][]string) int {
	var trees = 0
	x := 0
	y := 0
	for y < len(slopes) {
		if slopes[y][x] == "#" {
			trees += 1
		}
		y += down
		x = (x + right)%len(slopes[0])
	}
	fmt.Printf("(PART 2) Total count of trees: %d\n", trees)
	return trees
}