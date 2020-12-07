package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    file, err := os.Open("/Users/ahassoun/Documents/adventofcode/2020/day-1/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    // At first pass read all entries and store them
    var entries []int

    for scanner.Scan() {
        entry , err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        entries = append(entries, entry)
    }

    out:
    for i, entry1 := range entries {
        for y, entry2 := range entries {
            if i != y {
                sum := entry1 + entry2
                if sum == 2020 {
                    fmt.Printf("Found 2 entries that sum to 2020: %d and %d \n", entry1, entry2)
                    fmt.Printf("Multiplying them together: %d \n", entry1*entry2)
                    break out
                }
            }
        }
    }
}