package topic3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func operand() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	op := make([]int, 100)
	p := 0
	n := len(list)
	for i := 0; i < n; i++ {
		switch list[i] {
		case "+":
			op[p-2] += op[p-1]
			p--
		case "-":
			op[p-2] -= op[p-1]
			p--
		case "*":
			op[p-2] *= op[p-1]
			p--
		default:
			op[p], _ = strconv.Atoi(list[i])
			p++
		}
	}
	fmt.Println(op[0])
}
