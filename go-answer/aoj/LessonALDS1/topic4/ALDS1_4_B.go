package topic4

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

func binSearch(A []int, tcd int) bool {
	var first, last int
	last = len(A) - 1

	for first <= last {
		mid := (first + last) / 2
		if A[mid] == tcd {
			return true
		}
		if A[mid] < tcd {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return false
}

func main() {
	// Your code here!
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)

	var n, d int
	var S, T []int
	fmt.Sscan(getNextLine(scanner), &n)
	S = getIntList(scanner)
	fmt.Sscan(getNextLine(scanner), &d)
	T = getIntList(scanner)

	var count int
	for _, t := range T {
		if binSearch(S, t) {
			count++
		}
	}
	fmt.Println(count)
	writer.Flush()
}
