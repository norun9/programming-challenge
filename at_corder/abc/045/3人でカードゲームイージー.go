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

func judge(str string) string {
	if str == "a" {
		return "A"
	}
	if str == "b" {
		return "B"
	}
	return "C"
}

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)

	var A, B, C []string

	A = getStringList(scanner)
	B = getStringList(scanner)
	C = getStringList(scanner)

	turn := "A"
	for {
		if turn == "A" {
			//fmt.Println("A:", A)
			if len(A) == 0 {
				fmt.Println("A")
				break
			}
			turn = judge(A[0])
			A = A[1:]
			continue
		}
		if turn == "B" {
			//fmt.Println("B:", B)
			if len(B) == 0 {
				fmt.Println("B")
				break
			}
			turn = judge(B[0])
			B = B[1:]
			continue
		}
		if turn == "C" {
			//fmt.Println("C:", C)
			if len(C) == 0 {
				fmt.Println("C")
				break
			}
		}
		turn = judge(C[0])
		C = C[1:]
		continue
	}
	writer.Flush()
}
