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
	// var n, k int
	var S []int
	S = getIntList(scanner)
	y := S[0]
	x := S[1]
	list := make([][]int, y)
	for i := 0; i < y; i++ {
		A := getIntList(scanner)
		list[i] = A
	}

	for i := 0; i < y; i++ {
		var result []int
		for j := 0; j < x; j++ {
			nowPlace := list[i][j]
			count := nowPlace
			// 右
			if j < x-1 {
				for k := j + 1; k < x; k++ {
					count += list[i][k]
					// fmt.Println("右", list[i][k])
				}
			}
			// 左(現在位置x座標から左にマイナスしていく)
			if j > 0 {
				for k := j - 1; k >= 0; k-- {
					count += list[i][k]
					// fmt.Println("左", list[i][k])
				}
			}
			// 上(上にインデックスをマイナスしていく)
			if i > 0 {
				for k := i - 1; k >= 0; k-- {
					count += list[k][j]
					// fmt.Println("上", list[k][j])
				}
			}
			// 下
			if i < y-1 {
				for k := i + 1; k < y; k++ {
					count += list[k][j]
					// fmt.Println("下", list[k][j])
				}
			}
			result = append(result, count)
		}
		fmt.Println(strings.Trim(fmt.Sprintf("%v", result), "[]"))
	}
	writer.Flush()
}
