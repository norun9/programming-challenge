package _42

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

func find(A []int, tcd int) int {
	I := -1
	for i, a := range A {
		if a == tcd {
			I = i
		}
	}
	return I
}

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)

	var T []int
	T = getIntList(scanner)
	list := []int{5, 7, 5}
	l := len(list)
	count := 0
	for _, t := range T {
		result := find(list, t)
		if result == -1 {
			break
		} else {
			list = append(list[:result], list[result+1:]...)
			count++
		}
	}
	if count == l {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	writer.Flush()
}
