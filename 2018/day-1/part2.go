package main
import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strconv"
)

func main() {

  total, found, map_freqs, lines, err := readLines("input.txt")
  if err != nil {
    log.Fatalf("readLines: %s", err)
  }

  num := 0
  if ! found {
    for ! found {
      num += 1
      for i, freq := range lines {
        _ = i
        total += freq
        value, ok := map_freqs[total]
        if ok {
          map_freqs[total] = value + 1
          if map_freqs[total] == 2 {
            found = true
            break
          }
        } else {
          map_freqs[total] = 1
        }
      }
    }
    fmt.Printf("First duplicate freq: %d \n",total)
  }
  fmt.Printf("First duplicate freq: %d \n",total)
}

func readLines(path string) (int, bool, map[int]int, []int, error) {
  file, err := os.Open(path)
  var first = true
  var found = false
  total := 0

  if err != nil {
    log.Fatalf("readLines: %s", err)
  }
  defer file.Close()

  lines := make([]int, 0)
  map_freqs := make(map[int]int)

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    num, err := strconv.Atoi(scanner.Text())
    if err != nil {
      log.Fatalf("readLines: %s", err)
    }

    lines = append(lines, num)
    total += num

    if first {
      first = false
    } else {
      value, ok := map_freqs[total]
      if ok {
        map_freqs[total] = value + 1
        if map_freqs[total] == 2 {
          found = true
          break
        }
      } else {
        map_freqs[total] = 1
      }
    }
  }

  return total, found, map_freqs, lines, scanner.Err()
}
