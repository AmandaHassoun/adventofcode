package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// PART 1
	input := readFileIntoArray("input.txt")
	increases := numOfIncreases(input)
	log.Printf("Part1 answer: %d", increases)

	// PART 2
	increasesPart2 := windowMeasurements(input)
	log.Printf("Part2 answer: %d", increasesPart2)
}

func readFileIntoArray(filename string) []string {
	var input []string
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open input file")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	f.Close()

	return input
}

func numOfIncreases(depths []string) int {
	increases := 0

	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increases += 1
		}
	}

	//Check for the last item in the list
	if depths[len(depths)-1] > depths[len(depths)-2] {
		increases += 1
	}
	return increases
}

func windowMeasurements(depths []string) int {
	increases := 0
	letters := make(map[string]int)
	threeMeasurementsTracker := make(map[string]int)

	for i := 0; i < len(depths); i++ {
		window := strings.Split(depths[i], " ")
		unit, err := strconv.Atoi(window[0])
		if err != nil {
			log.Fatalf("failed to convert string unit to int")
		}
		for i := 1; i < len(window); i++ {
			if len(window[i]) > 0 {
				letters[window[i]] += unit
				threeMeasurementsTracker[window[i]] += 1
			}
		}
	}

	keys := make([]string, 0, len(letters))
	for k := range letters {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("map:", threeMeasurementsTracker)

	for i := 1; i < len(keys); i++ {
		if threeMeasurementsTracker[keys[i]] == 3 {
			if letters[keys[i]] > letters[keys[i-1]] {
				increases += 1
			}
		} else {
			break
		}
	}
	fmt.Println("keys:", keys)

	return increases
}
