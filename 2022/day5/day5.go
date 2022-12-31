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
    getCrateArrangement("input.txt", 1)
	topCrate := topOfStack(stackMap)
    fmt.Printf("(Part 1) Crates that end up on the top: %s \n \n", topCrate)

    // Reset global var
    stackMap = make(map[int][]string)

	// PART 2
	getCrateArrangement("input.txt", 2)
	topCrate = topOfStack(stackMap)
	fmt.Printf("(Part 2) Crates that end up on the top with the CrateMover 9001 : %s \n \n", topCrate)
}

func getCrateArrangement(filename string, part int) () {
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
            if part == 1 {
                moveFromStack9000(src, dest, count)
            } else if part == 2 {
                moveFromStack9001(src, dest, count)
            }
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

func moveFromStack9000(src int, dest int, count int) {
    for i := 0; i < count; i++ {
        srcStackCrate := stackMap[src][0]
        // Add to dest
        stackMap[dest] = append([]string{srcStackCrate}, stackMap[dest]...)
        //Remove from source
        stackMap[src] = append(stackMap[src][:0], stackMap[src][1:]...)
    }
}

func moveFromStack9001(src int, dest int, count int) {
    for i := count - 1; i >= 0; i-- {
        srcStackCrate := stackMap[src][i]
        // Add to dest
        stackMap[dest] = append([]string{srcStackCrate}, stackMap[dest]...)
        //Remove from source
        stackMap[src] = append(stackMap[src][:i], stackMap[src][i+1:]...)
    }
}

func topOfStack(stackMap map[int][]string) string {
    var topCrates string
    for i := 1; i <= len(stackMap); i++ {
        topCrates += stackMap[i][0]
    }

    return topCrates
}
