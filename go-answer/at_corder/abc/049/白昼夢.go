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

	var S string
	fmt.Sscan(getNextLine(scanner), &S)
	strByte := []byte(S)

	for len(strByte) > 0 {
		if len(strByte) >= 5 && (string(strByte[len(strByte)-5:]) == "dream" || string(strByte[len(strByte)-5:]) == "erase") {
			strByte = strByte[:len(strByte)-5]
		} else if len(strByte) >= 7 && string(strByte[len(strByte)-7:]) == "dreamer" {
			strByte = strByte[:len(strByte)-7]
		} else if len(strByte) >= 6 && string(strByte[len(strByte)-6:]) == "eraser" {
			strByte = strByte[:len(strByte)-6]
		} else {
			break
		}
	}

	if len(strByte) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	writer.Flush()
}
