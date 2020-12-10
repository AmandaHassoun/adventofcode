package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var allRules = make(map[string][]string)
var trackHits = make(map[string]int)

func main() {

	dir, err := os.Getwd()
	check(err)
	fileIO, err := os.OpenFile(dir+"/day-7/input.txt", os.O_RDWR, 0600)
	check(err)

	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {
		lineWithoutPeriod := strings.TrimSuffix(line, ".")
		parseRules(lineWithoutPeriod)
	}

	for k, v := range allRules {
		if !strings.Contains(k, "shiny gold bag") {
			getBagColorCount(k, v)
		}
	}

	fmt.Printf("Final count: %d \n", len(trackHits))
}

func parseRules(line string) {
	rules := strings.Split(line, "contain")
	trimTheS := strings.TrimSuffix(rules[0], "s ")
	strTrimSpace := strings.TrimSpace(trimTheS)
	for _, bagTypes := range rules[1:] {
		if strings.Contains(bagTypes, "no other bags") {
			allRules[strTrimSpace] = nil
		} else {
			remNumbers := strings.Split(bagTypes, ",")
			for _, bagType := range remNumbers {
				firstDigit := regexp.MustCompile(`[0-9]+`)
				bagList := firstDigit.Split(bagType, -1)
				rawBagType := strings.TrimSuffix(bagList[1], "s")
				nestedBagTrimSpace := strings.TrimSpace(rawBagType)
				allRules[strTrimSpace] = append(allRules[strTrimSpace], nestedBagTrimSpace)
			}
		}
	}
}

func getBagColorCount(key string, rules []string) {

	for _, nestedRules := range rules {
		if strings.Contains(nestedRules, "shiny gold bag") {
			_, exists := trackHits[key]
			if !exists {
				trackHits[key] = 1
			} else {
				trackHits[key]++
			}
		} else {
			getBagColorCount(key, allRules[nestedRules])
		}
	}
}
