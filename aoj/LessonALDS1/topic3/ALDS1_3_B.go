package topic3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Proc struct {
	name string
	time int
}

func queue() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	info := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(info[0])
	q, _ := strconv.Atoi(info[1])
	list := make([]Proc, n+1)
	for i := 0; i < n; i++ {
		sc.Scan()
		lists := strings.Split(sc.Text(), " ")
		time, _ := strconv.Atoi(lists[1])
		proc := Proc{name: lists[0], time: time}
		list[i] = proc
	}
	var first, last, count int
	first = 0
	last = n // 最後のインデックス＋１
	for first != last {
		done := q
		if done > list[first].time {
			done = list[first].time
		}
		count += done
		list[first].time -= done
		if list[first.time] == 0 {
			// countは少し考えないといけない
			fmt.Printf("%s %d\n", list[first].name, count)
		} else {
			list[last] = list[first]
			if last == n {
				last = 0
			} else {
				last++
			}
		}
		if first == n {
			first = 0
		} else {
			first++
		}
	}
}
