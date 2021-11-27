package topic6

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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
	return strings.Split(getNextLine(scanner), " ")
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

type Card struct {
	alpha string
	num   int
}

func quickSort(A []Card, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)
	}
}

func partition(A []Card, p, r int) int {
	x := A[r].num
	i := p - 1
	for j := p; j < r; j++ {
		if A[j].num <= x {
			i = i + 1
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

func mergeSort(A []Card) []Card {
	var result []Card
	if len(A) == 1 {
		return A
	}
	mid := len(A) / 2
	l := mergeSort(A[:mid])
	r := mergeSort(A[mid:])

	var i, j int
	for len(l) > i && len(r) > j {
		if l[i].num > r[j].num {
			result = append(result, r[j])
			j++
		} else {
			result = append(result, l[i])
			i++
		}
	}

	result = append(result, l[i:]...)
	result = append(result, r[j:]...)

	return result
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

	var cards []Card
	for i := 0; i < n; i++ {
		var A []string
		A = getStringList(scanner)
		alpha := A[0]
		x, _ := strconv.Atoi(A[1])
		cards = append(cards, Card{alpha: alpha, num: x})
	}
	t := make([]Card, len(cards))
	copy(t, cards)

	quickSort(cards, 0, n-1)

	mergeSorted := mergeSort(t)

	// 安定的なマージソートの配列と比較する.
	cond := reflect.DeepEqual(cards, mergeSorted)
	if cond {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}

	for _, v := range cards {
		fmt.Printf("%s %d\n", v.alpha, v.num)
	}
	writer.Flush()
}
