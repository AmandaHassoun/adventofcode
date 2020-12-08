package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dir, err := os.Getwd()
	check(err)
	file, err := os.Open(dir + "/day-3/input-part2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var slopes [][]string

	for scanner.Scan() {
		slope := strings.Split(scanner.Text(), "")
		slopes = append(slopes, slope)
	}

	slope1 := multiply(1,1, slopes)
	slope2 := multiply(3,1, slopes)
	slope3 :=multiply(5,1, slopes)
	slope4 := multiply(7,1, slopes)
	slope5 :=multiply(1,2, slopes)

	fmt.Printf("Total: %d\n", slope1*slope2*slope3*slope4*slope5)
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
	fmt.Printf("Total count of trees: %d\n", trees)
	return trees
}