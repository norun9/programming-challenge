package main

import (
	"bufio"
	"fmt"
	"math"
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

	var n int
	var A []int
	fmt.Sscan(getNextLine(scanner), &n)
	var pt, px, py int
	isValid := true
	for i := 0; i < n; i++ {
		A = getIntList(scanner)
		// d: 最低でも操作しなければいけない回数
		// 例: px=1, py=2, A[1]=1, A[2]=1
		// d = |1-1| + |2-1| = 1
		// 最低でも1回は操作しないと辿り着けない
		d := math.Abs(float64(px - A[1] + py - A[2]))
		dt := A[0] - pt
		// 移動可能回数(dt)よりも最低でも操作しなければいけない回数が上回ればNo
		// 到着したい時刻(T) < 最短到着時刻(X+Y)
		if dt < int(d) {
			isValid = false
			break
		}
		// 移動回数ー最低でも操作しなければいけない回数が偶数ならば実現可能
		// 座標への距離に対して移動距離が大きければ、それぞれの偶奇が一致していればOK、
		// そうでなければNG（偶奇が一致していれば、同じ場所で行ったりきたりして移動距離を調節できるため）
		if dt%2 != int(d)%2 {
			isValid = false
			break
		}
		pt = A[0]
		px = A[1]
		py = A[2]
	}
	if isValid {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	writer.Flush()
}
