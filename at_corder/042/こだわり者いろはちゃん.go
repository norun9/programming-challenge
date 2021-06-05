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

func getIntList(scanner *bufio.Reader) []int {
	list := strings.Split(getNextLine(scanner), " ")
	l := len(list)
	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i], _ = strconv.Atoi(list[i])
	}
	return result
}

func splitStr(A int) []string {
	return strings.Split(strconv.Itoa(A), "")
}

func isValid(n int, T []int) bool {
	l := len(strconv.Itoa(n))
	isValid := true
	for i := 0; i < l; i++ {
		// 1001なら1になる
		// 100なら0
		// 1004なら4
		tcd := n % 10
		// 1001なら次のループでは100
		n /= 10
		for _, t := range T {
			if t == tcd {
				isValid = false
				break
			}
		}
		if !isValid {
			break
		}
	}
	return isValid
}

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)

	var T, A []int
	T = getIntList(scanner)
	price := T[0]
	A = getIntList(scanner)

	for i := price; i < price*10; i++ {
		if isValid(i, A) {
			fmt.Println(i)
			return
		}
	}
	writer.Flush()
}
