package _44

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

func getUintList(scanner *bufio.Reader) []uint {
	list := strings.Split(getNextLine(scanner), " ")
	l := len(list)
	result := make([]uint, l)
	for i := 0; i < l; i++ {
		I, _ := strconv.Atoi(list[i])
		result[i] = uint(I)
	}
	return result
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

// 重複無し
func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)
	var N []int
	var A []uint
	N = getIntList(scanner)
	A = getUintList(scanner)

	n := uint(N[0])
	a := N[1]
	mp := map[uint]uint{}
	// n == 5
	// 00000 から 11111 までのフラグを管理していく
	// 例えばbit:00001の時
	// [00001, 00010, 00100, 01000, 10000]の中でAND演算に引っ掛かるのは00001のみであって
	// その時のindexは0である。
	// したがって[1 5 7 10 21]この中で1を探索する。
	// 2進数は00000, 00001, 00010,00011, 00100, 00101, 00111と数を増やしていき、
	// それぞれ[00001, 00010, 00100, 01000, 10000]の要素と一つずつAND演算子をすることで総当たりを実現できる
	// わからなくなれば→https://c-taquna.hatenablog.com/entry/2019/11/21/220511
	for bit := uint(0); bit < (1 << uint(n)); bit++ {
		var current uint = 0
		var count uint
		for i := uint(0); i < n; i++ {
			// 一つずつビットを左にずらしていく(2のn乗)
			// 例えば4回ループが回るなら[00001, 00010, 00100, 01000, 10000]となっていく
			if bit&(1<<i) != 0 {
				count++
				current += A[i]
			}
		}
		// mapにappendしていく
		mp[current] = count
	}
	var result int
	fmt.Println(mp)
	for sum, count := range mp {
		if count > 0 {
			if (float64(sum) / float64(count)) == float64(a) {
				fmt.Println("result:", float64(sum/count))
				fmt.Println("sum:", sum, "count:", count)
				result += 1
			}
		}
	}
	fmt.Println(result)
	writer.Flush()
}
