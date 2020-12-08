package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dir, err := os.Getwd()
	check(err)
	fileIO, err := os.OpenFile(dir+"/day-4/input-part2s.txt", os.O_RDWR, 0600)
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

func validBirthYear(byr string) bool {
	intBYR , err := strconv.Atoi(byr)
	check(err)
	if intBYR >= 1920 && intBYR <= 2002 {
		return true
	}
	return false
}

func validIssueYear(iyr string) bool {
	intIYR , err := strconv.Atoi(iyr)
	check(err)
	if intIYR >= 2010 && intIYR <= 2020 {
		return true
	}
	return false
}

func validExpirationYear(eyr string) bool {
	intEYR , err := strconv.Atoi(eyr)
	check(err)
	if intEYR >= 2020 && intEYR <= 2030 {
		return true
	}
	return false
}

func validHeight(height string) bool {
	if strings.Contains(height, "cm") {
		intHeight := getNumericalHeight(height, "cm")
		if intHeight >= 150 && intHeight <= 193 {
			return true
		}
	} else {
		intHeight := getNumericalHeight(height, "in")
		if intHeight >= 150 && intHeight <= 193 {
			return true
		}
	}
	return false
}

func getNumericalHeight(height string, unit string) int {
	onlyHeightnum := strings.ReplaceAll(height, unit, "")
	intHeight , err := strconv.Atoi(onlyHeightnum)
	check(err)
	return intHeight
}

func validHairColor(hc string) bool {
	var validContent = regexp.MustCompile(`^#[a-f0-9]+$`).MatchString
	length := len(hc[1:])
	if validContent(hc) && (length == 6){
		return true
	}
	return false
}

func validateEyeColor(ecl string) bool {
	validColors := [7]string{"amb", "blu", "brn" , "gry","grn", "hzl", "oth"}
	for _, color := range validColors {
		if strings.Contains(ecl, color) {
			return true
		}
	}
	return false
}

func validatePID(pid string) bool{
	if len(pid) == 9 {
		return true
	}
	return false
}
