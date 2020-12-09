package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var instructCountPart1 = make(map[string]int)
var instructCountPart2= make(map[string]int)
var accumulator int
var instructionsThatLoop []string
var program[]string
var firstLoop = true

func main() {

	dir, err := os.Getwd()
	check(err)
	fileIO, err := os.OpenFile(dir+"/day-8/input.txt", os.O_RDWR, 0600)
	check(err)

	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawBytes), "\n")
	for _, instr := range lines {
		program = append(program, instr)
	}
	//Part1
	runProgramPart1(program)

	//Part 2
	accumulator = 0
	processOfElimination(instructionsThatLoop)
}

func runProgramPart1(operations []string) {
	i := 0
out:
	for i < len(operations) {
		result, breakOut := returnOperationPart1(operations[i], i)
		if breakOut {
			break out
		}
		i += result
	}
}

func runProgramPart2(operations []string) bool {
	i := 0
	done := false
out:
	for i < len(operations) {
		result, breakOut := returnOperationPart2(operations[i], i)
		if breakOut {
			break out
		}
		i += result
		if (i == len(program)){
			done = true
			fmt.Printf("We done, accumulator is: %d \n", accumulator)
			break out
		}
	}
	return done
}

func returnOperationPart2(instruction string, ind int) (int, bool) {
	index := 1
	var breakOut = false
	var counts []string

	instructionsThatLoop = append(instructionsThatLoop, instruction)
	splitInstruction := strings.Split(instruction, " ")
	if strings.Contains(splitInstruction[1], "+") {
		counts = strings.SplitAfter(splitInstruction[1], "+")
	} else {
		counts = strings.SplitAfter(splitInstruction[1], "-")
	}
	num, err := strconv.Atoi(strings.TrimSpace(counts[1]))

	if err != nil {
		fmt.Println(err)
	}

	_, exists := instructCountPart2[instruction + strconv.Itoa(ind)]
	if exists {
		breakOut = true
	} else {
		instructCountPart2[instruction + strconv.Itoa(ind)] = 1
	}

	if splitInstruction[0] == "acc" {
		if counts[0] == "+" {
			accumulator += num
			index = 1
		} else {
			accumulator -= num
			index = 1
		}
	} else if splitInstruction[0] == "jmp" {
		if counts[0] == "+" {
			index = num
			jmpCount, _ := strconv.Atoi(counts[1])
			if jmpCount == 0 {
				index = 1
			}
		} else {
			index = -num
		}
	} else {
		index = 1
	}

	return index, breakOut
}

func processOfElimination(operations []string){
	var newString string
	out:
	for _, instr := range operations {
		if strings.Contains(instr, "jmp") {
			newString = strings.Replace(instr, "jmp", "nop", 1)
			newProgram := findAndReplace(instr, newString)
			accumulator = 0
			instructCountPart2 = make(map[string]int)
			done := runProgramPart2(newProgram)
			if done{
				break out
			}
		} else if strings.Contains(instr, "nop"){
			newString = strings.Replace(instr, "nop", "jmp", 1)
			newProgram := findAndReplace(instr, newString)
			accumulator = 0
			instructCountPart2 = make(map[string]int)
			done := runProgramPart2(newProgram)
			if done{
				break out
			}
		}
	}
}

func findAndReplace(instrToReplace string, replaceItwithMe string) []string {
	programsCopy := make([]string, len(program))
	_ = copy(programsCopy, program)
	for i, instr := range program {
		if instr == instrToReplace {
			programsCopy[i] = replaceItwithMe
		}
	}
	return programsCopy
}

func returnOperationPart1(instruction string, ind int) (int, bool) {
	index := 1
	var breakOut = false
	var counts []string

	instructionsThatLoop = append(instructionsThatLoop, instruction)
	splitInstruction := strings.Split(instruction, " ")
	if strings.Contains(splitInstruction[1], "+") {
		counts = strings.SplitAfter(splitInstruction[1], "+")
	} else {
		counts = strings.SplitAfter(splitInstruction[1], "-")
	}
	num, err := strconv.Atoi(strings.TrimSpace(counts[1]))

	if err != nil {
		fmt.Println(err)
	}

	_, exists := instructCountPart1[instruction + strconv.Itoa(ind)]
	if exists {
		if firstLoop {
			fmt.Printf("We hit the first infinite loop and acc is: %d \n", accumulator)
			firstLoop = false
		}
 		breakOut = true
	} else {
		instructCountPart1[instruction + strconv.Itoa(ind)] = 1
	}

	if splitInstruction[0] == "acc" {
		if counts[0] == "+" {
			accumulator += num
			index = 1
		} else {
			accumulator -= num
			index = 1
		}
	} else if splitInstruction[0] == "jmp" {
		if counts[0] == "+" {
			index = num
		} else {
			index = -num
		}
	} else {
		index = 1
	}

	return index, breakOut
}