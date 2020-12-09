package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	file, err := os.Open(dir + "/day-2/input.txt")
	check(err)
	defer file.Close()

	scannerPart1 := bufio.NewScanner(file)
	scannerPart1.Split(bufio.ScanLines)

	for scannerPart1.Scan() {
		part1(scannerPart1)
	}

	file, err = os.Open(dir + "/day-2/input.txt")
	check(err)
	defer file.Close()

	scannerPart2 := bufio.NewScanner(file)
	scannerPart2.Split(bufio.ScanLines)

	for scannerPart2.Scan() {
		part2(scannerPart2)
	}
}

func part1(s *bufio.Scanner) {
	validPasswords := 0
	for s.Scan() {
		line := strings.Split(s.Text(), ":")
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

func part2(s *bufio.Scanner){
	validPasswords := 0
	for s.Scan() {
		line := strings.Split(s.Text(), ":")
		policy := strings.TrimSpace(line[0])
		splitPolicy := strings.Fields(policy)
		validCounts := strings.Split(splitPolicy[0], "-")
		character := splitPolicy[1]
		password := strings.Split(strings.TrimSpace(line[1]), "")
		index1, err := strconv.Atoi(validCounts[0])
		check(err)
		index2, err := strconv.Atoi(validCounts[1])
		check(err)
		if (password[index1-1] == character || password[index2-1] == character) && !(password[index1-1] == character && password[index2-1] == character) {
			validPasswords += 1
		}
	}
	fmt.Printf("Total number of valid passwords: %d \n", validPasswords)
}