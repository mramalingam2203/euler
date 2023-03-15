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

func maxPrimeFactors(n uint64) uint64 {
	// Initialize the maximum prime factor
	// variable with the lowest one
	var maxPrime uint64 = 0

	// Prbig.Int the number of 2s that divide n
	for n%2 == 0 {
		maxPrime = 2
		n = n / 2 // equivalent to n /= 2
	}
	// n must be odd at this pobig.Int
	for n%3 == 0 {
		maxPrime = 3
		n = n / 3
	}

	// now we have to iterate only for big.Integers
	// who does not have prime factor 2 and 3
	for i := 5; uint64(i) <= uint64(math.Sqrt(float64(n))); i += 6 {
		for n%uint64(i) == 0 {
			maxPrime = uint64(i)
			n = n / uint64(i)
		}
		for n%(uint64(i)+2) == 0 {
			maxPrime = uint64(i + 2)
			n = n / uint64((i + 2))
		}
	}

	// This condition is to handle the case
	// when n is a prime number greater than 4
	if n > 4 {
		maxPrime = n
	}

	return maxPrime
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)
	if t < 1 || t > 10 {
		fmt.Println("Too many trials")
		os.Exit(0)
	}
	for tItr := 0; tItr < int(t); tItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		if n < 1 || n > 1e12 {
			fmt.Println("n off range")
			os.Exit(0)
		}
		fmt.Println(maxPrimeFactors(uint64(n)))

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
