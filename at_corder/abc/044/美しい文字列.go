package _44

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNextLine(scanner *bufio.Reader) string {
	var buffer string
	for {
		line, isPrefix, _ := scanner.ReadLine()
		buffer += string(line)
		if isPrefix == false {
			break
		}
	}
	return buffer
}

func getIntList(scanner *bufio.Reader) []int {
	list := strings.Split(getNextLine(scanner), " ")
	l := len(list)
	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i], _ = strconv.Atoi(list[i])
	}
	return result
}

type CountNum struct {
	Str string
	Num int
}

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)

	var w string
	fmt.Sscan(getNextLine(scanner), &w)
	list := strings.Split(w, "")
	m := map[string]int{}
	for _, l := range list {
		_, ok := m[l]
		if ok {
			m[l] += 1
		} else {
			m[l] = 1
		}
	}

	for _, value := range m {
		if value%2 != 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
	writer.Flush()
}
