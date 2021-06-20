package main

import (
	"bufio"
	"fmt"
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

func getStringList(scanner *bufio.Reader) []string {
	return strings.Split(getNextLine(scanner), "")
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)
	var A []int
	A = getIntList(scanner)
	W := A[0]
	H := A[1]
	N := A[2]
	X, Y := 0, 0
	for i := 0; i < N; i++ {
		var T []int
		T = getIntList(scanner)
		x := T[0]
		y := T[1]
		n := T[2]

		if n == 1 {
			X = max(X, x)
		}

		if n == 2 {
			W = min(W, x)
		}

		if n == 3 {
			Y = max(Y, y)
		}

		if n == 4 {
			H = min(H, y)
		}
	}
	rst := max(W-X, 0) * max(H-Y, 0)
	fmt.Println(rst)

	writer.Flush()
}
