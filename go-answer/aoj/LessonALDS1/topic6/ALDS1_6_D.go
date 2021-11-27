package topic6

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Min(integers ...int) int {
	m := integers[0]
	for i, integer := range integers {
		if i == 0 {
			continue
		}
		if m > integer {
			m = integer
		}
	}
	return m
}

type node struct {
	value int
	place int
}

func solve(m int, seq []int) int {
	cost := 0

	// initialize
	temp := make([]node, m)
	for i, s := range seq {
		temp[i].value = s
		temp[i].place = i
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].value < temp[j].value
	})

	nodes := make([]node, m)
	for i, s := range seq {
		nodes[i].value = s
		nodes[temp[i].place].place = i
	}

	s := temp[0].value

	// compute
	for i := 0; i < m; i++ {
		j := nodes[i].place
		if j >= 0 && j != i {
			n := 1

			amin := nodes[i].value
			sum := nodes[i].value
			for j != i {
				next := nodes[j].place
				if nodes[j].value < amin {
					amin = nodes[j].value
				}
				sum += nodes[j].value
				n++
				nodes[j].place = -1
				j = next
			}
			cost += min(sum+(n-2)*amin, sum+amin+(n+1)*s)
		}
	}
	return cost
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
	var wList []int
	wList = getIntList(scanner)

	ans := solve(n, wList)
	fmt.Println(ans)
	writer.Flush()
}
