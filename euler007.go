package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)
	if t < 1 || t > 1000 {
		fmt.Println("Too many tests")
		os.Exit(0)
	}
	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		if n < 1 || n > 1e4 {
			fmt.Println("Number out of range")
			os.Exit(0)
		}
		fmt.Println(countPrimes(int(n)))
	}
}


func countPrimes(limit int) int {
	primes := make([]int, limit)
	count := 0
	isPrimeDivisible := func(candidate int) bool {
		for j := 0; j < count; j++ {
			if math.Sqrt(float64(candidate)) < float64(primes[j]) {
				return false
			}
			isDivisibe := isDivisible(candidate, primes[j])
			if isDivisibe {
				return true
			}
		}
		return false
	}
	for candidate := 2; ; {
		if count < limit {
			if !isPrimeDivisible(candidate) {
				primes[count] = candidate
				count++
			}
			candidate++
		} else {
			break
		}
	}
	return primes[limit-1]
}

func isDivisible(candidate, by int) bool {
	return candidate%by == 0
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
