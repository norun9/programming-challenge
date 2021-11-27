package main

import (
	"bufio"
	"fmt"
	"os"
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

func getStrList(scanner *bufio.Reader) []string {
	return strings.Split(getNextLine(scanner), "")
}

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)

	var A []string
	A = getStrList(scanner)
	var stack []string
	for _, a := range A {
		switch a {
		case "0":
			stack = append(stack, "0")
		case "1":
			stack = append(stack, "1")
		case "B":
			if len(stack) > 0 {
				stack = append([]string{}, stack[:len(stack)-1]...)
			}
		}
	}
	fmt.Println(strings.Join(stack, ""))
	writer.Flush()
}
