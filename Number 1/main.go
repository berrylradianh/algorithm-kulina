package main

import (
	"fmt"
	"sort"
)

func sockMerchant(n int, ar []int) int {
	sort.Ints(ar)
	pairCount := 0
	i := 0

	for i < n-1 {
		if ar[i] == ar[i+1] {
			pairCount++
			i += 2
		} else {
			i++
		}
	}

	return pairCount
}

func main() {
	var n int
	fmt.Scan(&n)

	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	result := sockMerchant(n, ar)
	fmt.Println(result)
}
