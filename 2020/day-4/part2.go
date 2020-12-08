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
	fileIO, err := os.OpenFile(dir+"/day-4/input-part2.txt", os.O_RDWR, 0600)
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
		properPassport := strings.ReplaceAll(passport, "\n", " ")
		for _, field := range required {
			if strings.Contains(properPassport, field+":") {
				fieldCount += 1
			}
		}
		if fieldCount == 8 {
			validFieldsCount := validatePassport(strings.Split(properPassport, " "))
			if validFieldsCount == 7 {
				validPassports += 1
			}
		}

		if fieldCount == 7 && !(strings.Contains(properPassport, "cid:")) {
			validFieldsCount := validatePassport(strings.Split(properPassport, " "))
			if validFieldsCount == 7 {
				validPassports += 1
			}
		}
		fieldCount = 0
	}
	fmt.Printf("Total # of valid passports: %d\n", validPassports)
}

func validatePassport(passport []string) int {
	validationCount := 0
	for _, field := range passport {
		if strings.Contains(field, "byr"){
			isBYRvalid := validBirthYear(strings.Split(field, ":")[1])
			if isBYRvalid{
				validationCount += 1
			}
		}
		if strings.Contains(field, "iyr"){
			isIYRvalid := validIssueYear(strings.Split(field, ":")[1])
			if isIYRvalid{
				validationCount += 1
			}
		}
		if strings.Contains(field, "eyr"){
			isEYRvalid := validExpirationYear(strings.Split(field, ":")[1])
			if isEYRvalid{
				validationCount += 1
			}
		}
		if strings.Contains(field, "hgt"){
			isHGTvalid := validHeight(strings.Split(field, ":")[1])
			if isHGTvalid{
				validationCount += 1
			}
		}
		if strings.Contains(field, "hcl"){
			isHCLvalid := validHairColor(strings.Split(field, ":")[1])
			if isHCLvalid{
				validationCount += 1
			}
		}
		if strings.Contains(field, "ecl"){
			isECLvalid := validateEyeColor(strings.Split(field, ":")[1])
			if isECLvalid{
				validationCount += 1
			}
		}
		if strings.Contains(field, "pid"){
			isPIDvalid := validatePID(strings.Split(field, ":")[1])
			if isPIDvalid{
				validationCount += 1
			}
		}
	}
	return validationCount
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
		if intHeight >= 59 && intHeight <= 76 {
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
