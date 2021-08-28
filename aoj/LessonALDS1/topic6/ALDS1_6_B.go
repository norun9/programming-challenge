package topic6

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

func partition(A []int, p, r int) int {
	x := A[r] // 基準値

	i := p - 1
	var t int
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			t = A[i]
			A[i] = A[j]
			A[j] = t
		}
	}
	t = A[i+1]
	A[i+1] = A[r]
	A[r] = t
	return i + 1
}

func partitionOfMine(A []int, p, r int) ([]int, int) {
	x := A[r] // 基準

	var j int
	for idx, a := range A {
		if a <= x {
			prev := A[j]
			A[j] = a
			A[idx] = prev
			j++
		}
	}
	return A, j - 1
}

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Sscan(getNextLine(scanner), &n)

	var A []int
	A = getIntList(scanner)

	A, ans := partitionOfMine(A, 0, len(A)-1)

	var result []string
	for i, a := range A {
		if i == ans {
			result = append(result, fmt.Sprintf("[%d]", a))
		} else {
			result = append(result, fmt.Sprintf("%d", a))
		}
	}
	fmt.Println(strings.Join(result, " "))
	writer.Flush()
}
