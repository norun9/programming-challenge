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
	list := strings.Split(getNextLine(scanner), "")
	l := len(list)
	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i], _ = strconv.Atoi(list[i])
	}
	return result
}

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)
	A := getNextLine(scanner)
	k := len(A)
	result := 0
	for bit := uint(0); bit < 1<<uint(k-1); bit++ {
		fmt.Println("**************************")
		fmt.Printf("bit: %08b\n", bit)
		var st int
		for i := uint(0); i < uint(k-1); i++ {
			if bit&(1<<i) != 0 {
				fmt.Printf("index bit: %08b\n", 1<<i)
				num, _ := strconv.Atoi(A[st : int(i)+1])
				fmt.Println("st:", st)
				fmt.Println("i+1:", i+1)
				fmt.Println("num:", num)
				result += num
				st = int(i) + 1
			}
		}
		if st < k {
			num, _ := strconv.Atoi(A[st:k])
			fmt.Println("in case st < k num:", num)
			result += num
		}
	}

	fmt.Println("******** result *********")
	fmt.Printf("%d\n", result)
	writer.Flush()
}
