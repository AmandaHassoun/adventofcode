package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// PART 1
	totalScore := getStrategyGuideScore("input.txt")
	fmt.Printf("(Part 1) Total score: %d \n \n", totalScore)
}

func getStrategyGuideScore(filename string) (int) {
    totalScore := 0
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open input file!!")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
	    strategy := strings.Split(scanner.Text(), " ")
	    totalScore += getScore(strategy)
	}

	f.Close()

	return totalScore
}

func getScore(round []string) (int) {
    //Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
    // A for Rock, B for Paper, and C for Scissors
    // X for Rock, Y for Paper, and Z for Scissors
    // 1 for Rock, 2 for Paper, and 3 for Scissors)
    // 0 if you lost, 3 if the round was a draw, and 6 if you won

    score := 0
    if round[0] == "A" {
        if round[1] == "Y" {
            score = 2 + 6
        } else if round[1] == "Z" {
            score = 0 + 3
        } else {
            score = 3 + 1
        }
    } else if round[0] == "B" {
        if round[1] == "Y" {
            score = 3 + 2
        } else if round[1] == "Z" {
            score = 6 + 3
        } else {
            score = 0 + 1
        }
    } else {
        if round[1] == "Y" {
            score = 0 + 2
        } else if round[1] == "Z" {
            score = 3 + 3
        } else {
            score = 6 + 1
        }
    }

    return score
}
