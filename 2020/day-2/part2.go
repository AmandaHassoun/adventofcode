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

	//part1()
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
