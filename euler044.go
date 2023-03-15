package main

import (
	"fmt"
	"os"
)

func pentagonalNumber(n int32) []int32 {
	numbers := make([]int32, 0)
	var i int32 = 1

	for ; ; i++ {
		numbers = append(numbers, int32(i*(3*i-1)/2))
		if numbers[len(numbers)-1] >= n {
			break
		}
	}
	return numbers
}

func main() {
	// no. of trials and number
	var K, N int32
	fmt.Scanf("%d %d", &N, &K)

	if K < 1 || K > 9999 {
		fmt.Println(" K off range")
		os.Exit(0)
	}

	if N < K+1 || N > 1e6 {
		fmt.Println(" N off range")
		os.Exit(0)
	}

}
