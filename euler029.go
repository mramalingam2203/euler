//https://www.hackerrank.com/contests/projecteuler/challenges/euler029/problem?isFullScreen=true

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var N uint64
	fmt.Scanf("%d", &N)
	if N < 2 || N > 1e5 {
		fmt.Println("Number out of range")
		os.Exit(0)
	}

	list := []float64{}
	var i, j uint64
	for i = 2; i <= N; i++ {
		for j = 2; j <= N; j++ {
			//fmt.Println(math.Log(float64(j)) * float64(i))
			list = append(list, float64(math.Log(float64(j))*float64(i)))
		}
	}
	//sort.uint64(list)
	fmt.Print(len(removeDuplicateValues(list)))

}

func removeDuplicateValues(float64Slice []float64) []float64 {
	keys := make(map[uint64]bool)
	list := []float64{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range float64Slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
