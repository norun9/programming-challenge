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

func splitStr(A int) []string {
	return strings.Split(strconv.Itoa(A), "")
}

func getStrList(scanner *bufio.Reader) []string {
	return strings.Split(getNextLine(scanner), "")
}

func issyo() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	var A []int
	fmt.Sscan(getNextLine(scanner), &n)
	A = getIntList(scanner)
	var count int
	for _, a := range A {
		count += a
	}
	avg := math.Ceil(float64(count) / float64(len(A)))
	avgFloor := math.Floor(float64(count) / float64(len(A)))

	var sum int
	var sum2 int
	for _, a := range A {
		sum2 += int(math.Pow(float64(a-int(avgFloor)), 2))
		sum += int(math.Pow(float64(a-int(avg)), 2))
	}
	if sum > sum2 {
		fmt.Println(sum2)
	} else {
		fmt.Println(sum)
	}
	writer.Flush()
}
