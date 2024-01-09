package main

import (
	"bufio"
	"fmt"
	"os"
)

func countingValleys(n int, s string) int {
	altitude := 0
	valleyCount := 0

	for _, step := range s {
		if step == 'U' {
			altitude++
		} else if step == 'D' {
			altitude--
		}

		if altitude == 0 && step == 'U' {
			valleyCount++
		}
	}

	return valleyCount
}

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	result := countingValleys(n, s)

	fmt.Println(result)
}
