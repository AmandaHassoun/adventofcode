package main
import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    file, err := os.Open("input1.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
	num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
	}
        count += num
    }

    fmt.Printf("Total: %d \n", count)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
