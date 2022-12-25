package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// PART 1
	maxCalories, caloriesSumList := findMaxCaloriesSum("input.txt")
	fmt.Printf("(Part 1) Elf carrying the most Calories total: %d \n \n", maxCalories)

	// PART 2
	topThreeSum := findTopThreeSum(caloriesSumList)
	fmt.Printf("(Part 2) Sum of Top three Elves carrying the most Calories: %d \n \n", topThreeSum)
}

func findMaxCaloriesSum(filename string) (int, []int) {
    var caloriesSumList []int
	var currentMax int
	totalCaloriesSum := 0
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open input file!!")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

    currentMax = 0
	for scanner.Scan() {
	    if len(scanner.Text()) != 0 {
            calories, _ := strconv.Atoi(scanner.Text())
            totalCaloriesSum += calories
            continue
	    }

        caloriesSumList = append(caloriesSumList, totalCaloriesSum)

	    if totalCaloriesSum > currentMax {
	        currentMax = totalCaloriesSum
	    }
	    totalCaloriesSum = 0
	}

	f.Close()

	return currentMax, caloriesSumList
}

func findTopThreeSum(calorieSums []int) (int) {
    topThree := []int{0,0,0}
    topThree[0] = calorieSums[0]
    topThree[1] = calorieSums[1]
    topThree[2] = calorieSums[2]

    currentMinIndex, currentMinVal := findMin(topThree)
    for i := 3; i < len(calorieSums); i++ {
	    if calorieSums[i] > currentMinVal {
            topThree[currentMinIndex] = calorieSums[i]
	    }
	    currentMinIndex, currentMinVal = findMin(topThree)
    }

    topThreeSum := topThree[0] + topThree[1] + topThree[2]

	return topThreeSum
}

func findMin(someList []int) (int, int) {
    var min int = someList[0]
    var index int

    for i, value := range someList {
        if min > value {
            min = value
            index = i
        }
    }

    return index, min
}
