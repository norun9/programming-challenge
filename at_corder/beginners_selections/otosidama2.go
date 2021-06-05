package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// 	"math"
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

func otosidama2() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)
	var A []int
	A = getIntList(scanner)
	// 	N := A[0]
	Y := A[1]
	var man, gosen, sen int
	for Y != 0 {
		if Y > 10000 {
			man += Y / 10000
			Y -= 10000 * man
		}
		if Y > 5000 {
			gosen += Y / 5000
			Y -= 5000 * gosen
		}
		if Y > 1000 {
			sen += Y / 1000
			Y -= 1000 * sen
		} else {
			fmt.Println("can't")
			break
		}
	}
	fmt.Println(man, gosen, sen)
	writer.Flush()
}
