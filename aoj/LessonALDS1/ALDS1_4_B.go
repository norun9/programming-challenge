package LessonALDS1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binSearch(tcd int, A []string) int {
	low := 0
	high := len(A) - 1
	for low <= high {
		mid := (low + high) / 2
		a, _ := strconv.Atoi(A[mid])
		guess := a
		if guess == tcd {
			return mid
		}
		if guess > tcd {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func bin() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	sc.Scan()
	S := strings.Split(sc.Text(), " ")

	sc.Scan()
	sc.Scan()
	T := strings.Split(sc.Text(), " ")
	var count int
	for _, t := range T {
		tt, _ := strconv.Atoi(t)
		var i int
		if i = binSearch(tt, S); i >= 0 {
			count++
		}
	}
	fmt.Println(count)
}
