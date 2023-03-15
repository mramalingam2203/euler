package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const N int32 = 1000

var dp [N + 1]int32

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
		ch := make(chan int) // Create a new channel.
		go Generate(ch)      // Launch Generate goroutine.
		sum := 0
		for i := 0; i < 4; i++ {
			prime := <-ch
			sum += prime
			print(sum, "\n")
			ch1 := make(chan int)
			go Filter(ch, ch1, prime)
			ch = ch1
		}
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

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}
