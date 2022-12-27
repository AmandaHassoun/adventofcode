package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	// PART 1
	containedRangesCount, _, overlapPairs := getFullyContainedRangeCount("input.txt")
	fmt.Printf("(Part 1) Total count of assignment pairs where one range fully contains the other: %d \n \n", containedRangesCount)

	// PART 2
	fmt.Printf("(Part 2) Total count of assignment pairs where ranges overlap: %d \n \n", overlapPairs)
}

func getFullyContainedRangeCount(filename string) (int, int, int) {
    count := 0
    numPairs := 0
    overlaps := 0
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open input file!!")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
	    strategy := strings.Split(scanner.Text(), ",")
	    contained, num := checkRanges(strategy[0], strategy[1])
	    checkOverlap := checkOverlap(strategy[0], strategy[1])
        if contained {
            count++
            numPairs += num
        }
        if checkOverlap {
            overlaps++
        }
	}

	f.Close()

	return count, numPairs, overlaps
}

func checkRanges(range1 string, range2 string) (bool, int) {
    var numericInterval1, numericInterval2 [2]int
    interval1 := strings.Split(range1, "-")
    interval2 := strings.Split(range2, "-")

    for i := 0; i < len(interval1); i++ {
        numericInterval1[i], _ = strconv.Atoi(interval1[i])
        numericInterval2[i], _ = strconv.Atoi(interval2[i])
    }

    if (numericInterval1[0] >= numericInterval2[0]) && (numericInterval1[1] <= numericInterval2[1]) {
        return true, (numericInterval1[1] - numericInterval1[0]) + 1
    } else if (numericInterval2[0] >= numericInterval1[0]) && (numericInterval2[1] <= numericInterval1[1]) {
        return true, (numericInterval2[1] - numericInterval2[0]) + 1
    }

    return false, 0
}

func checkOverlap(range1 string, range2 string) (bool) {
    var numericInterval1, numericInterval2 [2]int
    interval1 := strings.Split(range1, "-")
    interval2 := strings.Split(range2, "-")

    for i := 0; i < len(interval1); i++ {
        numericInterval1[i], _ = strconv.Atoi(interval1[i])
        numericInterval2[i], _ = strconv.Atoi(interval2[i])
    }

    if (numericInterval1[0] >= numericInterval2[0]) && (numericInterval1[1] <= numericInterval2[1]) {
        return true
    } else if (numericInterval2[0] >= numericInterval1[0]) && (numericInterval2[1] <= numericInterval1[1]) {
        return true
    } else if (numericInterval1[0] <= numericInterval2[1]) && (numericInterval1[0] >= numericInterval2[0]) {
        return true
    } else if (numericInterval1[1] >= numericInterval2[0]) && (numericInterval1[0] <= numericInterval2[1]) {
        return true
    }

    return false
}
