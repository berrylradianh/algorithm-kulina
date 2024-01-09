package main

import "fmt"

func main() {
	inputNumber := 1345679

	for divisor := 1000000; divisor >= 1; divisor /= 10 {
		digit := (inputNumber / divisor) % 10
		fmt.Println(digit * divisor)
	}
}
