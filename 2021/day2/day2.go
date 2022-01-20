package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// PART 1
	input := readFileIntoArray("input.txt")
	finalPosition := getPosition(input)
	log.Printf("Part1 answer: %d", finalPosition)

	//PART 2
	finalPositionWithAim := getPositionWithAim(input)
	log.Printf("Part2 answer: %d", finalPositionWithAim)
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

func getPosition(course []string) int {
	var depth, horizontal int

	for i := 0; i < len(course); i++ {
		position := strings.Split(course[i], " ")
		unit, err := strconv.Atoi(position[1])
		if err != nil {
			log.Fatalf("failed to convert string unit to int")
		}
		if position[0] == "forward" {
			horizontal += unit
		} else if position[0] == "down" {
			depth += unit
		} else if position[0] == "up" {
			depth -= unit
		}
	}

	return depth * horizontal
}

func getPositionWithAim(course []string) int {
	var depth, horizontal, aim int

	for i := 0; i < len(course); i++ {
		position := strings.Split(course[i], " ")
		unit, err := strconv.Atoi(position[1])
		if err != nil {
			log.Fatalf("failed to convert string unit to int")
		}
		if position[0] == "forward" {
			horizontal += unit
			depth = depth + (unit * aim)
		} else if position[0] == "down" {
			aim += unit
		} else if position[0] == "up" {
			aim -= unit
		}
	}

	return depth * horizontal
}
