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

var mask int
var mem = make(map[int]int)
var setPositions []int
var unsetPositions []int
var decimalTotal int

func main() {

	dir, err := os.Getwd()
	check(err)
	fileIO, err := os.OpenFile(dir+"/day-14/input.txt", os.O_RDWR, 0600)
	check(err)

	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {
		lineSplit := strings.Split(line, "=")
		trimMask := strings.TrimSpace(lineSplit[1])
		if strings.TrimSpace(lineSplit[0]) == "mask" {
			setPositions = returnPositionsToSet(trimMask)
			unsetPositions = returnPositionsToUnset(trimMask)
		} else {
			memIndex := getMemLocation(strings.TrimSpace(lineSplit[0]))
			val, err := strconv.Atoi(strings.TrimSpace(lineSplit[1]))
			check(err)
			for i := 0; i < len(unsetPositions); i++ {
				val = clearBit(val, uint(unsetPositions[i]))
			}
			for i := 0; i < len(setPositions); i++ {
				val = setBit(val, uint(setPositions[i]))
			}
			mem[memIndex] = val
		}
	}

	for _, v := range mem {
		decimalTotal += v
	}
	//Part 1
	fmt.Println(decimalTotal)
}

func returnPositionsToSet(mask string) []int {
	var setPositions []int
	for i := 0; i < len(mask); i++ {
		if mask[i] == '1' {
			setPositions = append(setPositions, 35-i)
		}
	}
	return setPositions
}

func returnPositionsToUnset(mask string) []int {
	var unsetPositions []int
	for i := 0; i < len(mask); i++ {
		if mask[i] == '0' {
			unsetPositions = append(unsetPositions, 35-i)
		}
	}
	return unsetPositions
}

func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func clearBit(n int, pos uint) int {
	return n &^ (1 << pos)
}

func getMemLocation(memory string) int {
	splitMem := strings.Split(memory, "[")
	getNum := strings.Split(splitMem[1], "]")
	intVal, err := strconv.Atoi(strings.TrimSpace(getNum[0]))
	check(err)
	return intVal
}
