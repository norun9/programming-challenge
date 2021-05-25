package LessonALDS1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type StackArea struct {
	popIndex int
	area     int
}

func sumArea() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	infos := strings.Split(sc.Text(), "")
	stack := make([]int, 0)
	stackArea := make([]StackArea, 0)
	stackArea = append(stackArea, StackArea{popIndex: 0, area: 0})
	for i, info := range infos {
		if info == "\\" {
			stack = append(stack, i)
		}
		if info == "/" {
			// pop
			if len(stack) > 0 {
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stackArea) > 0 {
					stackAreaLast := stackArea[len(stackArea)-1].popIndex
					area := i - last
					// popしたindexが、stackAreaに保存したタプルに入っているindexより大きい時
					if last > stackAreaLast {
						stackArea = append(stackArea, StackArea{popIndex: last, area: area})
					} else {
						// 合計した水溜りは人繋がりなので一つのたぷるに格納したい
						// popしたindexが、stackAreaに保存したタプルに入っているindexより小さい時
						lastArea := stackArea[len(stackArea)-1].area
						// popする
						stackArea = stackArea[:len(stackArea)-1]
						sumArea := lastArea + area
						stackArea = append(stackArea, StackArea{popIndex: last, area: sumArea})
					}
				}
			}
		}
	}
	sum := 0
	for _, s := range stackArea {
		sum += s.area
	}
	fmt.Println(sum)
}
