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
var invalidNum int

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
		sumCombos := findSumCombos(numbersToSkip)
		inList := sumInList(allNumbers[i], sumCombos)
		if !inList {
			fmt.Printf("This number doesn't add up to previous: %d\n ",allNumbers[i])
			invalidNum = allNumbers[i]
			break out
		}
	}

	var currSum = 0
	var contiguousNums []int
	var initialIndex = 0

	for i, _ := range allNumbers {
		for currSum < invalidNum {
			currSum += allNumbers[i]
			contiguousNums = append(contiguousNums, allNumbers[i])
			i++
		}
		if currSum == invalidNum {
			min, max := findMinMax(contiguousNums)
			fmt.Printf("Sum of min&max: %d \n", min+max)
			break
		} else{
			i = initialIndex + 1
			currSum = 0
			contiguousNums = nil
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

func findMinMax(sumList []int) (int,int) {
	var max int = sumList[0]
	var min int = sumList[0]
	for _, num := range sumList {
		if max < num {
			max = num
		}
		if min > num {
			min = num
		}
	}
	return min, max
}