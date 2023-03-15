// https://www.hackerrank.com/contests/projecteuler/challenges/euler006/problem?isFullScreen=true

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
		fmt.Println("too many tests expected")
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
		result(n)
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

func result(N int32) {
	var n uint64
	n = uint64(N)
	fmt.Println((n * n * (n + 1) * (n + 1) / 4) - (n * (n + 1) * (2*n + 1) / 6))

}
