package main

import "fmt"

func main() {
	switches := make([]bool, 100)

	for trip := 1; trip <= 100; trip++ {
		for switchNumber := 1; switchNumber <= 100; switchNumber++ {
			if switchNumber%trip == 0 {
				switches[switchNumber-1] = !switches[switchNumber-1]
			}
		}
	}

	lampsOn := 0
	for _, state := range switches {
		if state {
			lampsOn++
		}
	}

	fmt.Printf("After 100 trips, %d lamps will be turned on.\n", lampsOn)
}
