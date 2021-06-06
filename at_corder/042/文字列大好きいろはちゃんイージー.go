package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func getStrList(scanner *bufio.Reader) []string {
	list := strings.Split(getNextLine(scanner), "")
	return list
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
	N := T[0]
	strList := []string{}
	for i := 0; i < N; i++ {
		S := getNextLine(scanner)
		strList = append(strList, S)
	}
	sort.Slice(strList, func(i, j int) bool { return strList[i] < strList[j] })
	fmt.Println(strings.Join(strList, ""))
	writer.Flush()
}
