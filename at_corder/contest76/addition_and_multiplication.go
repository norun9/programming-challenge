package at_coder

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
	// Your code here!
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)
	var n, k int
	fmt.Sscan(getNextLine(scanner), &n)
	fmt.Sscan(getNextLine(scanner), &k)
	now := 1
	for i := 0; i < n; i++ {
		addTmp := now + k
		mulTmp := now * 2
		if addTmp <= mulTmp {
			now += addTmp
		} else {
			now += mulTmp
		}
	}
	fmt.Println(now)
	writer.Flush()
}
