package topic2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectionSort() {
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
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if ints[min] > ints[j] {
				min = j
			}
		}
		if min != i {
			count++
			ints[i], ints[min] = ints[min], ints[i]
		}
	}
	fmt.Println(strings.Trim(fmt.Sprintf("%v", ints), "[]"))
	fmt.Println(count)
}
