package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getNextLine(scanner *bufio.Reader) string {
	var buffer string
	for {
		line, isPrefix, _ := scanner.ReadLine()
		buffer += string(line)
		if isPrefix == false {
			break
		}
	}
	return buffer
}
func getIntList(scanner *bufio.Reader) []int {
	list := strings.Split(getNextLine(scanner), " ")
	l := len(list)
	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i], _ = strconv.Atoi(list[i])
	}
	return result
}

func main() {
	// Your code here!
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)
	var N, K int
	l := getIntList(scanner)
	N = l[0]
	K = l[1]
	var A, B []int
	A = getIntList(scanner)
	B = getIntList(scanner)
	var diff float64
	for i := 0; i < N; i++ {
		diff += math.Abs(float64(A[i]) - float64(B[i]))
	}
	if int(diff) > K {
		fmt.Println("No")
		return
	}

	if int(diff)%2 != K%2 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	writer.Flush()
}
