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

func getIntList(scanner *bufio.Reader) []int {
	list := strings.Split(getNextLine(scanner), " ")
	l := len(list)
	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i], _ = strconv.Atoi(list[i])
	}
	return result
}

// https://qiita.com/drken/items/fd4e5e3630d0f5859067#%E7%AC%AC-8-%E5%95%8F--abc-085-c---otoshidama-300-%E7%82%B9
func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)
	var A []int
	A = getIntList(scanner)
	N := A[0]
	Y := A[1]
	// 一万円札の枚数
	for i := 0; i < N+1; i++ {
		// 五千円札の枚数
		for j := 0; j < N-i+1; j++ {
			// 一千円札の枚数
			k := N - i - j
			// 全ての合計が合計金額と一致する場合
			if 10000*i+5000*j+1000*k == Y {
				fmt.Println(i, j, N-i-j)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
	writer.Flush()
}
