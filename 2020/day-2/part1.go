package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {

	dir, err := os.Getwd()
	check(err)
	file, err := os.Open(dir + "/day-2/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	validPasswords := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		policy := strings.TrimSpace(line[0])
		splitPolicy := strings.Fields(policy)
		validCounts := strings.Split(splitPolicy[0], "-")
		character := splitPolicy[1]
		password := strings.TrimSpace(line[1])
		actualCount := strings.Count(password, character)
		minCount, err := strconv.Atoi(validCounts[0])
		check(err)
		maxCount, err := strconv.Atoi(validCounts[1])
		check(err)
		if actualCount >= minCount && actualCount <= maxCount {
			validPasswords += 1
		}
	}
	fmt.Printf("Total number of valid passwords: %d \n", validPasswords)
}
