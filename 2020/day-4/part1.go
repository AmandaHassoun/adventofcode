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
	fileIO, err := os.OpenFile(dir+"/day-4/input-part1.txt", os.O_RDWR, 0600)
	check(err)

	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}
	required := [8]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	fieldCount := 0
	validPassports := 0

	lines := strings.Split(string(rawBytes), "\n\n")
	for _, passport := range lines {
		for _, field := range required {
			if strings.Contains(passport, field+":") {
				fieldCount += 1
			}
		}
		if fieldCount == 8 {
			validPassports += 1
		}

		if fieldCount == 7 && !(strings.Contains(passport, "cid:")) {
			validPassports += 1
		}
		fieldCount = 0
	}
	fmt.Printf("Total # of valid passports: %d\n", validPassports)
}
