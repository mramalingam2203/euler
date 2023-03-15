package main

import (
	"fmt"
	"os"
)

func triangleNumber(n int32) []int32 {
	numbers := make([]int32, 0)
	var i int32 = 1
	for ; ; i++ {
		numbers = append(numbers, int32(i*(i+1)/2))
		if numbers[len(numbers)-1] >= n {
			break
		}
	}
	return numbers

}

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

func hexagonalNumber(n int32) []int32 {
	numbers := make([]int32, 0)
	var i int32 = 1
	for ; ; i++ {
		numbers = append(numbers, int32(i*(2*i-1)))
		if numbers[len(numbers)-1] >= n {
			break
		}
	}
	return numbers
}

func intersection(a []int32, b []int32) []int32 {
	var inBoth []int32

	m := make(map[int32]int32)
	for _, k := range a {
		m[k] |= (1 << 0)
	}
	for _, k := range b {
		m[k] |= (1 << 1)
	}

	for k, v := range m {
		a := v&(1<<0) != 0
		b := v&(1<<1) != 0
		if a && b {
			inBoth = append(inBoth, k)
		}
	}
	return inBoth
}

func main() {
	//fmt.Println(triangleNumber(10000))
	//fmt.Println(pentagonalNumber(10000))
	//fmt.Println(hexagonalNumber(10000))
	var N, a, b int32
	fmt.Scanf("%d %d %d", &N, &a, &b)
	//Number := [3]int{3, 5, 6}
	if a == b {
		fmt.Println("Identical filters..cant proceed")
		os.Exit(0)
	}
	/*
		if a != 3 || a != 5 || a != 6 || b != 3 || b != 5 || b != 6 {
			fmt.Println("Irrelevant indices")
			os.Exit(0)

		}*/

	if a == 3 && b == 5 || a == 5 && b == 3 {
		 := intersection(triangleNumber(N), pentagonalNumber(N))
	} else if a == 5 && b == 6 || a == 6 && b == 5 {
		array := intersection(hexagonalNumber(N), pentagonalNumber(N))
	} else if a == 3 && b == 6 || a == 6 && b == 3 {
		array := intersection(triangleNumber(N), hexagonalNumber(N))
	}
	fmt.Print(array)
}
