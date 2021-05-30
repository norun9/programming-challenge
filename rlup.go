package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
func convInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func main() {
	lineInfos := strings.Split(nextLine(), " ")
	h, _ := strconv.Atoi(lineInfos[0])
	units := make([][]string, 0)
	for i := 0; i < h; i++ {
		unit := strings.Split(nextLine(), " ")
		units = append(units, unit)
	}
	dx := []int{1, 0, -1, 0}
	dy := []int{0, -1, 0, 1}
	R, U, L, D := 0, 1, 2, 3
	var count, x, y int
	for i := 0; i <= 4; i++ {
		if i == 0 || i == 2 {
			// Rigth
			x += dx[R]
			y += dy[R]
			count += convInt(units[y][x])
		}
		if i == 1 {
			// Down
			x += dx[D]
			y += dy[D]
			count += convInt(units[y][x])
		}
		if i == 3 {
			// Up
			x += dx[U]
			y += dy[U]
			count += convInt(units[y][x])
		}
		if i == 4 {
			// Left
			x += dx[L]
			y += dy[L]
			count += convInt(units[y][x])
		}
	}
	fmt.Println(count) // 17
}
