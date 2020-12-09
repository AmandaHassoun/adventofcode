package main

import (
	"bufio"
	"os"
	"strings"
)



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


}

