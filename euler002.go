package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Returns sum of even Fibonacci numbers which are
// less than or equal to given limit.
func evenFibSum(limit uint64) uint64 {
	if limit < 2 {
		return 0
	}

	// Initialize first two even  Fibonacci  numbers
	// and their sum
	var ef1, ef2, sum uint64
	ef1, ef2 = 0, 2
	sum = ef1 + ef2

	// calculating sum of even Fibonacci value
	for ef2 <= limit {
		// get next even value of Fibonacci sequence
		var ef3 uint64
		ef3 = 4*ef2 + ef1

		// If we go beyond limit, we break loop
		if ef3 > limit {
			break
		}

		// Move to next even number and update sum
		ef1 = ef2
		ef2 = ef3
		sum += ef2
	}

	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)
	if t < 1 || t > 1e5 {
		fmt.Println("TOo many trials")
		os.Exit(0)
	}

	for tItr := 0; tItr < int(t); tItr++ {
		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		if n < 10 || n > 4e16 {
			fmt.Println("Invalid range")
			os.Exit(0)
		}
		fmt.Println(evenFibSum(uint64(n)))
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
