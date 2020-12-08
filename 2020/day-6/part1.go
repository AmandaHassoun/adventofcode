package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	dir, err := os.Getwd()
	check(err)
	fileIO, err := os.OpenFile(dir+"/day-6/input-part1.txt", os.O_RDWR, 0600)
	check(err)

	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}

	var yesCount = 0
	lines := strings.Split(string(rawBytes), "\n\n")
	for _, group := range lines {
		answers := strings.Split(group, "\n")
		yesCount += countUniqueAnswers(answers)
	}
	fmt.Printf("Total # of yesses: %d\n", yesCount)
}

func countUniqueAnswers(answers []string) int {
	var answerCount map[string]int
	answerCount = make(map[string]int)
	for _, answer := range answers {
		personAnswers := strings.Split(answer, "")
		for _, question := range personAnswers {
			answerCount[question] = 1
		}
	}
	return len(answerCount)
}