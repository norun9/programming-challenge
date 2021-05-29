package aoj

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}


// 素数判定:
// 調査した値の倍数は全て除去する → その数の倍数はその数で割り切れる為、合成数となる
func isPrime(x int) bool {
	if x == 2 {
		return true
	}
	if x < 2 || x%2 == 0 {
		return false
	}
	i := 3
	for float64(i) <= math.Sqrt(float64(x)) {
		if x%i == 0 {
			return false
		}
		i += 2
	}
	return true
}

func main() {
	num, _ := strconv.Atoi(nextLine())
	var count int
	for i := 0; i < num; i++ {
		x, _ := strconv.Atoi(nextLine())
		if isPrime(x) {
			count += 1
		}
	}
	fmt.Println(count)
}
