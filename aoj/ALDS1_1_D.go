package aoj

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func maxProfit() {
	var n, r, min int
	// 最大の利益が負になることを考慮して、初期値は 10^9×(−1)以下に設定
	// -2147483648
	max := math.MinInt32
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	min, _ = strconv.Atoi(sc.Text())
	for i := 1; i < n; i++ {
		sc.Scan()
		r, _ = strconv.Atoi(sc.Text())
		result := r - min
		if result > max {
			max = result
		}
		if r < min {
			min = r
		}
	}
	fmt.Println(max)
}
