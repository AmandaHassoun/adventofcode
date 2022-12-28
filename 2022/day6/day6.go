package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// PART 1
	startOfPacket := parseFile("input.txt")
	fmt.Printf("(Part 1) Start of packet after: %d chars \n \n", startOfPacket)

	// PART 2
}

func parseFile(filename string) int {
	f, err := os.Open(filename)
	var index int

	if err != nil {
		log.Fatalf("Failed to open input file!!")
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
	    index = findStartOfPacket(scanner.Text())
	}

	f.Close()

	return index
}

func findStartOfPacket(datastream string) (int){
    var packet []string
    var numOfChars int

    for index , v := range datastream {

		if !contains(packet, string(v)) {
		    packet = append(packet, string(v))
		    if (len(packet) == 4) && checkForDuplicates(packet) {
                packet = append(packet[:0], packet[1:]...)
		    } else {
		        if (len(packet) == 4) {
		            numOfChars = index + 1
		            break
		        }
		    }
		} else {
		    // Remove first element from the list and add new one
		    packet = append(packet[:0], packet[1:]...)
		    packet = append(packet, string(v))
		}
	}

    return numOfChars
}

func contains(s []string, str string) (bool) {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func checkForDuplicates(charList []string) bool {
   visited := make(map[string]bool, 0)

   for i := 0; i < len(charList); i++{
      if visited[charList[i]] == true{
         return true
      } else {
         visited[charList[i]] = true
      }
   }
   return false
}
