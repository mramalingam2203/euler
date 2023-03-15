// What is the value of the first triangle number to have over  divisors?

package main

import (
	"fmt"
	"math"
)

func numDivisors(triangle int32) int32 {
	var factors int32 = 0
	for i := 1; i <= int(math.Pow(float64(triangle), 0.5))+1; i++ {
		if int(triangle)%i == 0 {
			factors += 1
		}
	}
	return factors * 2
}

func maxTriangleDivisors(max int32) int32 {
	var i int32 = 1
	var triangle int32 = 0
	if i > 0 {
		triangle += 1
		if numDivisors(triangle) >= max {
			return triangle
		}
		i += 1
	}
	return 0
}

func main() {
	fmt.Println(maxTriangleDivisors(500))
}
