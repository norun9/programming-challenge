package topic5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Int() uint {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return uint(i)
}

func main() {
	sc.Split(bufio.ScanWords)

	var n = Int()
	A := []uint{}
	for i := uint(0); i < n; i++ {
		A = append(A, Int())
	}

	var q = Int()
	mi := []uint{}
	for i := uint(0); i < q; i++ {
		mi = append(mi, Int())
	}

	mp := map[uint]bool{}
	// n == 5
	// 00000 から 11111 までのフラグを管理していく
	// 例えばbit:00001の時
	// [00001, 00010, 00100, 01000, 10000]の中でAND演算に引っ掛かるのは00001のみであって
	// その時のindexは0である。
	// したがって[1 5 7 10 21]この中で1を探索する。
	// 2進数は00000, 00001, 00010,00011, 00100, 00101, 00111と数を増やしていき、
	// それぞれ[00001, 00010, 00100, 01000, 10000]の要素と一つずつAND演算子をすることで総当たりを実現できる
	// わからなくなれば→https://c-taquna.hatenablog.com/entry/2019/11/21/220511
	for bit := uint(0); bit < (1 << n); bit++ {
		var current uint = 0
		fmt.Printf("現在のBit: %05b\n", bit)
		for i := uint(0); i < n; i++ {
			fmt.Printf("index: %d\n", i)
			fmt.Printf("(1<<i): %05b\n", (1 << i))
			// 一つずつビットを左にずらしていく(2のn乗)
			// 例えば4回ループが回るなら[00001, 00010, 00100, 01000, 10000]となっていく
			if bit&(1<<i) != 0 {
				fmt.Printf("AND演算子で引っかかった: %d\n", A[i])
				current += A[i]
			}
		}
		// mapにappendしていく
		fmt.Println("current:", current)
		mp[current] = true
	}
	for _, m := range mi {
		// mapでの存在チェックは非常に速い
		if _, ok := mp[m]; ok {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
