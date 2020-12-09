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

var instructCount = make(map[string]int)
var accumulator int

func main() {

	dir, err := os.Getwd()
	check(err)
	fileIO, err := os.OpenFile(dir+"/day-8/input-part1.txt", os.O_RDWR, 0600)
	check(err)

	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}

	var program[]string
	lines := strings.Split(string(rawBytes), "\n")
	for _, instr := range lines {
		program = append(program, instr)
	}
	runProgram(program)
}

func runProgram(operations []string){
	i := 0
	out:
	for i < len(operations) {
		result, breakOut:= returnOperation(operations[i], i)
		if breakOut {
			break out
		}
		i += result
	}
}

func returnOperation(instruction string, ind int) (int, bool) {
	index := 1
	var breakOut = false
	var counts []string

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

	_, exists := instructCount[instruction + strconv.Itoa(ind)]
	if exists {
 		fmt.Printf("Accumulator before we restart the loop is: %d \n", accumulator)
 		breakOut = true
	} else {
		instructCount[instruction + strconv.Itoa(ind)] = 1
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