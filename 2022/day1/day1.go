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
	maxCalories := findMaxCaloriesSum("input.txt")
	fmt.Printf("Elf carrying the most Calories total: %d", maxCalories)
}

func findMaxCaloriesSum(filename string) int {
	var current_max int
	totalCaloriesSum := 0
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open input file!!")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

    current_max = 0
	for scanner.Scan() {
	    if len(scanner.Text()) != 0 {
            calories, _ := strconv.Atoi(scanner.Text())
            totalCaloriesSum += calories
            continue
	    }

	    if totalCaloriesSum > current_max {
	        current_max = totalCaloriesSum
	    }
	    totalCaloriesSum = 0
	}

	f.Close()

	return current_max
}