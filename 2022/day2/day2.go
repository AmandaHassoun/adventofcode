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
	totalScorePart1, totalScorePart2 := getStrategyGuideScore("input.txt")
	fmt.Printf("(Part 1) Total score: %d \n \n", totalScorePart1)

	// PART 2
	fmt.Printf("(Part 2) Total score: %d \n \n", totalScorePart2)
}

func getStrategyGuideScore(filename string) (int, int) {
    totalScorePart1 := 0
    totalScorePart2 := 0
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open input file!!")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
	    strategy := strings.Split(scanner.Text(), " ")
	    totalScorePart1 += getScore(strategy)
	    totalScorePart2 += getScorePart2(strategy)
	}

	f.Close()

	return totalScorePart1, totalScorePart2
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

func getScorePart2(round []string) (int) {
    // X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
    // Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
    // A for Rock, B for Paper, and C for Scissors
    // X for Rock, Y for Paper, and Z for Scissors
    // 1 for Rock, 2 for Paper, and 3 for Scissors)
    // 0 if you lost, 3 if the round was a draw, and 6 if you won

    score := 0
    if round[0] == "A" {
        if round[1] == "Y" {
            score = 1 + 3
        } else if round[1] == "Z" {
            score = 2 + 6
        } else {
            score = 3 + 0
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
            score = 3 + 3
        } else if round[1] == "Z" {
            score = 6 + 1
        } else {
            score = 2 + 0
        }
    }

    return score
}
