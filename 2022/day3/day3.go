package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	// PART 1
	sum:= getSumOfPriorities("input.txt")
	fmt.Printf("(Part 1) Sum of the priorities of item types: %d \n \n", sum)

	// PART 2
	sumPart2 := getSumOfPrioritiesPart2("input.txt")
	fmt.Printf("(Part 2) Sum of the priorities of item types: %d \n \n", sumPart2)
}

func getSumOfPriorities(filename string) (int) {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open input file!!")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	priorities := getPriorities()

    sum := 0
	for scanner.Scan() {
	    strategy := strings.Split(scanner.Text(), "")
	    compartment1 := strategy[0:(len(strategy)/2)]
	    compartment2 := strategy[(len(strategy)/2):]
	    commonElement := findCommonElementTwo(compartment1, compartment2)
	    sum += priorities[commonElement]
	}

	f.Close()

	return sum
}

func getSumOfPrioritiesPart2(filename string) (int) {
    var setOfThree [][]string
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open input file!!")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	priorities := getPriorities()

    count := 0
    sum := 0
	for scanner.Scan() {
	    if count == 3 {
	        commonElement := findCommonElementThree(setOfThree[0], setOfThree[1], setOfThree[2])
	        sum += priorities[commonElement]
	        setOfThree = nil
	        count = 0
	    }
	    strategy := strings.Split(scanner.Text(), "")
	    setOfThree = append(setOfThree, strategy)
	    count++
	}

	commonElement := findCommonElementThree(setOfThree[0], setOfThree[1], setOfThree[2])
	sum += priorities[commonElement]

	f.Close()

	return sum
}

func getPriorities() (map[string]int) {
    priorities := make(map[string]int)

    index := 1
    for letter := 'a'; letter <= 'z'; letter++ {
		uppercaseLetter := unicode.ToUpper(letter)
		priorities[string(letter)] = index
		priorities[string(uppercaseLetter)] = index + 26
		index++
	}

	return priorities
}

func findCommonElementTwo(compartment1 []string, compartment2 []string) (string) {

    for i := 0; i < len(compartment1); i++ {
        for x := 0; x < len(compartment2); x++ {
		    if compartment1[i] == compartment2[x] {
		        return compartment1[i]
		    }
		}
	}
	return ""
}

func findCommonElementThree(compartment1 []string, compartment2 []string, compartment3 []string) (string) {

    for i := 0; i < len(compartment1); i++ {
        for x := 0; x < len(compartment2); x++ {
            for y := 0; y < len(compartment3); y++ {
		        if (compartment1[i] == compartment2[x]) && (compartment1[i] == compartment3[y]) && (compartment2[x] == compartment3[y]) {
		            return compartment1[i]
		        }
		    }
		}
	}
	return ""
}