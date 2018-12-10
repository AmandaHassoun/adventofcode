package main
import (
    "bufio"
    "fmt"
    "os"
    "log"
)


func main() {

    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scan := bufio.NewScanner(file)
    code_ids := make([]string, 0)

    for scan.Scan() {
      line := scan.Text()
      code_ids = append(code_ids, line)
    }

    if err := scan.Err(); err != nil {
        log.Fatal(err)
    }

    final := find_match(code_ids)
    fmt.Printf("Best match: %s \n", final)
}

func find_match(ids []string) (string) {

  least_diff := 30
  var match string

  for i, id1 := range ids {
    _ = id1
    for j, id2 := range ids {
      _ = id2
      if i != j {
        num_diffs, match_str := compare_strings(ids[i], ids[j])
        if num_diffs < least_diff {
          least_diff = num_diffs
          match = match_str
        }
      }
    }
  }
  return match
}

func compare_strings(id1 string, id2 string) (int, string) {

  diff := 0
  var match string

  for i, char := range id1 {
    _ = char
    if id1[i] != id2[i] {
      diff += 1
    }else{
      match += string(id1[i])
    }
  }
  return diff, match
}
