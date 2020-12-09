package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var instructCount = make(map[string]int)
var preamble = 25
var numsToConsider = 25

func main() {

	dir, err := os.Getwd()
	check(err)
	fileIO, err := os.OpenFile(dir+"/day-9/input-part1.txt", os.O_RDWR, 0600)
	check(err)

	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}

	var allNumbers[]int

	lines := strings.Split(string(rawBytes), "\n")
	for _, num := range lines {
		numInt, err := strconv.Atoi(num)
		check(err)
		allNumbers = append(allNumbers, numInt)
	}
	out:
	for i, _ := range allNumbers {
		i += preamble
		numbersToSkip := allNumbers[i-numsToConsider:i]
		fmt.Println(numbersToSkip)
		sumCombos := findSumCombos(numbersToSkip)
		inList := sumInList(allNumbers[i], sumCombos)
		if inList {
			fmt.Println(allNumbers[i])
		} else {
			fmt.Printf("This number doesn't add up to previous: %d\n ",allNumbers[i])
			break out
		}
	}
}

func findSumCombos(numbers []int) []int {
	var combos []int
	for i, _ := range numbers {
		for y := i+1; y < len(numbers); y++  {
			combos = append(combos, numbers[i]+numbers[y])
		}
	}
	return combos
}

func sumInList(num int, sumList []int) bool {
	for _, sum := range sumList {
		if sum == num {
			return true
		}
	}
	return false
}