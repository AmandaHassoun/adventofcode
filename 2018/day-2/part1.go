package main
import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strings"
)


func main() {

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    threes := 0
    total_threes := 0
    var found_threes bool = false
    twos := 0
    var found_twos bool = false
    total_twos := 0

    for scanner.Scan() {
      line := scanner.Text()
      for i, c := range line {
        num_of_char := strings.Count(line, string(c))
        _ = i
        if num_of_char == 2 && !found_twos {
          twos = 1
          found_twos = true
        }
        if num_of_char == 3 && !found_threes {
          threes = 1
          found_threes = true
        }
      }
      total_threes += threes
      total_twos += twos
      twos = 0
      threes = 0
      found_threes = false
      found_twos = false
    }
    fmt.Printf("checksum = %d \n", total_threes * total_twos)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
