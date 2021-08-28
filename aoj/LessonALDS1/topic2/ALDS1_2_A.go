package topic2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	strs := strings.Split(sc.Text(), " ")
	var ints []int
	for _, in := range strs {
		i, _ := strconv.Atoi(in)
		ints = append(ints, i)
	}
	var count int
	flag := true
	for flag {
		flag = false
		for i := n - 1; 0 < i; i-- {
			if ints[i] < ints[i-1] {
				ints[i], ints[i-1] = ints[i-1], ints[i]
				count++
				flag = true
			}
		}
	}
	fmt.Println(strings.Trim(fmt.Sprintf("%v", ints), "[]"))
	fmt.Println(count)
}
