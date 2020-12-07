package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dir, err := os.Getwd()
	check(err)
	file, err := os.Open(dir + "/day-1/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// At first pass read all entries and store them
	var entries []int

	for scanner.Scan() {
		entry, err := strconv.Atoi(scanner.Text())
		check(err)
		entries = append(entries, entry)
	}

out:
	for i, entry1 := range entries {
		for y, entry2 := range entries {
			for x, entry3 := range entries {
				if i != y && y != x && i != x {
					sum := entry1 + entry2 + entry3
					if sum == 2020 {
						fmt.Printf("Found 3 entries that sum to 2020: %d, %d and %d \n", entry1, entry2, entry3)
						fmt.Printf("Multiplying them together: %d \n", entry1*entry2*entry3)
						break out
					}
				}
			}
		}
	}
}
