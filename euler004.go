package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func largest_palindrome(n int32) int32 {
	fmt.Println(reverse(strconv.Itoa(int(n))))
	return 100

}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	fmt.Println(result)
	return
}

func is_Palindrome(n int32) bool {
	number := strconv.Itoa(int(n))
	reversed_number := reverse(number)
	fmt.Print(reversed_number)
	return false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)
	if t < 1 || t > 100 {
		fmt.Println("No.of trials out of range")
		os.Exit(0)
	}
	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		if n < 101101 || n > 1000000 {
			fmt.Println("Number  out of range")
			os.Exit(0)
		}
		_ = n
		is_Palindrome(n)

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
