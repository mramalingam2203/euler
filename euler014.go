//https://www.hackerrank.com/contests/projecteuler/challenges/euler014/problem?isFullScreen=true

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
	if t < 1 || t > 1e4 {
		fmt.Println("No.of trials out of range")
		os.Exit(0)
	}
	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		_ = n

		if n < 1 || n > 1e8 {
			fmt.Println("No.of trials out of range")
			os.Exit(0)
		}
		var LIMIT int32 = 10000
		var iteration_count, max int32 = 0, 0
		var i, count int32 = 0, 0
		_ = count
		for i = 1; i < LIMIT; i++ {
			iteration_count = count_iterations(int32(i))
			if iteration_count > max {
				max = iteration_count
				count = i
			}

		}

		fmt.Println(max)

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

func iteration(value int32) int32 {
	if value%2 == 0 {
		return value / 2
	} else {
		return 3*value + 1
	}
}

func count_iterations(value int32) int32 {
	var count int32 = 1

	if value != 1 {
		value = iteration(value)
		count += 1

	}
	return count
}
