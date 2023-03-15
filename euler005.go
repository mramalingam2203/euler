//https://www.hackerrank.com/contests/projecteuler/challenges/euler005/problem?isFullScreen=true

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)
	if t < 1 || t > 10 {
		fmt.Println("Too many tests to run")
		os.Exit(0)
	}

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		if n < 1 || n > 40 {
			fmt.Println("Number out of range")
			os.Exit(0)
		}
		_ = n
		lcm(n)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Function returns the lcm of first n numbers
func lcm(n int32) {
	var ans uint64 = 1
	var i uint64
	for i = 1; i <= uint64(n); i++ {
		ans = (ans * i) / (GCDEuclidean(ans, i))
	}
	fmt.Println(ans)
}

// GCDEuclidean calculates GCD by Euclidian algorithm.
func GCDEuclidean(a, b uint64) uint64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
