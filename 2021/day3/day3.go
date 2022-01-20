package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// PART 1
	input := readFileIntoArray("input.txt")
	counts := findCounts(input)
	gamma, epsilon := getRates(counts)
	log.Printf("Part1 answer: %d", gamma*epsilon)

	//PART 2
	getO2Rating(input, 0)
	//getCO2Rating(input, 0)

	//getRating(input, scrubber)
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

func findCounts(binary []string) map[int][]int {
	var occurrences = make(map[int][]int)

	for i := 0; i < len(binary); i++ {
		position := strings.Split(binary[i], "")
		for y := 0; y < len(position); y++ {
			if _, ok := occurrences[y]; ok {
				if position[y] == "0" {
					occurrences[y][0] += 1
				} else {
					occurrences[y][1] += 1
				}
			} else {
				occurrences[y] = []int{0, 0}
			}
		}

	}
	return occurrences
}

func getRates(binary map[int][]int) (int64, int64) {
	var gamma, epsilon string

	for i := 0; i < len(binary); i++ {
		if binary[i][0] > binary[i][1] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaBase64, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonBase64, _ := strconv.ParseInt(epsilon, 2, 64)

	return gammaBase64, epsilonBase64
}

func getO2Rating(input []string, index int) []string {
	var matches []string
	counts := findCounts(input)
	y := index

	if y == len(input[0]) {
		return matches
	} else {
		for i := 0; i < len(input); i++ {
			if counts[y][0] > counts[y][1] {
				if string(input[i][y]) == "0" {
					matches = append(matches, input[i])
				}
			} else if counts[y][0] <= counts[y][1] {
				if string(input[i][y]) == "1" {
					matches = append(matches, input[i])
				}
			}
			fmt.Println(matches)
		}
		y += 1
		getO2Rating(matches, y)
	}
	return matches
}

func getCO2Rating(input []string, index int) []string {
	var matches []string
	counts := findCounts(input)
	y := index

	if y == len(input[0]) {
		return matches
	} else {
		for i := 0; i < len(input); i++ {
			if counts[y][0] > counts[y][1] {
				if string(input[i][y]) == "1" {
					matches = append(matches, input[i])
				}
			} else if counts[y][0] <= counts[y][1] {
				if string(input[i][y]) == "0" {
					matches = append(matches, input[i])
				}
			}
			fmt.Println(matches)
		}
		y += 1
		getCO2Rating(matches, y)
	}
	return matches
}
