package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

var stackMap = make(map[int][]string)

func main() {
	// PART 1
    getCrateArrangement("input.txt")
	topCrate := topOfStack(stackMap)
    fmt.Printf("(Part 1) Crates that end up on the top: %v\n \n", topCrate)

	// PART 2
	//fmt.Printf("(Part 2) : %d \n \n", overlapPairs)
}

func getCrateArrangement(filename string) () {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open input file!!")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
	    line := scanner.Text()
	    if (!strings.Contains(line, "move") && !strings.Contains(line, "1")) {
            addToStackMap(line)
	    } else if strings.Contains(line, "move") {
            move := strings.Split(line, " ")
            src, _ := strconv.Atoi(move[3])
            dest, _ := strconv.Atoi(move[5])
            count, _ := strconv.Atoi(move[1])
            moveFromStack(src, dest, count)
	    }
	}

	f.Close()
}

func addToStackMap(stack string) {
    stackIndex := 1

    for charIndex, _ := range stack {
        if (charIndex <= len(stack) && (charIndex % 4) == 1) {
            if (stack[charIndex] > 'A' || stack[charIndex] < 'Z') && (stack[charIndex] != ' ') {
                stackMap[stackIndex] = append(stackMap[stackIndex], string(stack[charIndex]))
                charIndex += 2
            }
            stackIndex++
        }
    }
}

func moveFromStack(src int, dest int, count int) {
    for i := 0; i < count; i++ {
        srcStackCrate := stackMap[src][0]
        // Add to dest
        stackMap[dest] = append([]string{srcStackCrate}, stackMap[dest]...)
        //Remove from source
        stackMap[src] = append(stackMap[src][:0], stackMap[src][1:]...)
    }
}

func topOfStack(stackMap map[int][]string) string {
    var topCrates string
    for i := 1; i <= len(stackMap); i++ {
        topCrates += stackMap[i][0]
    }

    return topCrates
}
