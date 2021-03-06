package _44

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

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)

	var N, K, X, Y int
	fmt.Sscan(getNextLine(scanner), &N)
	fmt.Sscan(getNextLine(scanner), &K)
	fmt.Sscan(getNextLine(scanner), &X)
	fmt.Sscan(getNextLine(scanner), &Y)

	var sum int
	for i := 1; i <= N; i++ {
		if i <= K {
			sum += X
		} else {
			sum += Y
		}
	}
	fmt.Println(sum)
	writer.Flush()
}
