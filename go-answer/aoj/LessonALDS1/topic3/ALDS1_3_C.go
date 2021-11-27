package topic3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFirstPlace(A []int, key int) int {
	for i, a := range A {
		if a == key {
			return i
		}
	}
	return -1
}

func unset(A []int, place int) []int {
	if place >= len(A) {
		return A
	}
	return append(A[:place], A[place+1:]...)
}

func LessonALDS1() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	list := make([]int, 0)
	// var first, last int
	for i := 0; i < n; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		cmd := line[0]
		var key int
		if len(line) > 1 {
			key, _ = strconv.Atoi(line[1])
		}

		if cmd == "inert" {
			tmps := list[:]
			slice := []int{key}
			list = append(slice, tmps...)
		} else if cmd == "delete" {
			var place int
			if place = getFirstPlace(list, key); 0 <= place {
				list = unset(list, place)
			}
		} else if cmd == "deleteFirst" {
			list = list[1:]
		} else {
			list = list[:len(list)-1]
		}
	}

	fmt.Println(strings.Trim(fmt.Sprintf("%v", list), "[]"))
}
