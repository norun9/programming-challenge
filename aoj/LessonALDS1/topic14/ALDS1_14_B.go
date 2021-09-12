package topic14

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

	var T, P []string
	T = getStringList(scanner)
	P = getStringList(scanner)

	table := make(map[string]int)

	var isDupLast bool
	lastStr := P[len(P)-1]
	for i, j := len(P)-1, 0; i >= 0; i-- {
		if j != len(P)-1 && P[j] == lastStr {
			isDupLast = true
		}
		table[P[j]] = i
		j++
	}
	if !isDupLast {
		table[lastStr] = len(P)
	}
	//fmt.Println(table)

	i, p := len(P)-1, 0

	for i < len(T) {
		p = len(P) - 1

		for p >= 0 && i < len(T) {
			if T[i] == P[p] {
				i--
				p--
			} else {
				break
			}
		}
		if p < 0 {
			fmt.Println(i + 1)
		}

		var shift1, shift2 int
		if i > 0 {
			if shift, ok := table[T[i]]; !ok {
				shift1 = len(P)
			} else {
				shift1 = shift
			}
		}
		shift2 = len(P) - p
		i += max(shift1, shift2)
	}

	writer.Flush()
}
