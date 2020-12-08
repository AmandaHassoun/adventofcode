package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	dir, err := os.Getwd()
	check(err)
	file, err := os.Open(dir + "/day-3/input-part1.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var slopes [][]string

	for scanner.Scan() {
		slope := strings.Split(scanner.Text(), "")
		slopes = append(slopes, slope)
	}

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
	fmt.Printf("Total count of trees: %d\n", trees)
}
