package LessonALDS1

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

func getStringList(scanner *bufio.Reader) []string {
	return strings.Split(getNextLine(scanner), "")
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Tree struct {
	parent int
	left   int
	right  int
}

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Sscan(getNextLine(scanner), &n)

	trees := make([]Tree, n)
	depth := make([]int, n)
	for i := 0; i < n; i++ {
		trees[i].parent = -1
		trees[i].right = -1
		trees[i].left = -1
	}

	for i := 0; i < n-1; i++ {
		var A []int
		A = getIntList(scanner)
		id := A[0]
		k := A[1]
		l := -1
		if k > 0 {
			for j, v := range A[2:] {
				fmt.Println(v)
				trees[v].parent = id
				if j == 0 {
					trees[id].left = v
				} else {
					trees[l].right = v
				}
				l = v
			}
		}
	}
	for i, v := range trees {
		fmt.Println("index:", i)
		fmt.Println("parent:", v.parent)
		fmt.Println("right:", v.right)
		fmt.Println("left:", v.left)
		fmt.Println("***************")
	}
	for i, tree := range trees {
		if tree.parent == -1 {
			setDepth(trees, depth, i, 0)
		}
	}
	fmt.Println(depth)
	writer.Flush()
}

func setDepth(trees []Tree, depth []int, i int, d int) {
	depth[i] = d
	if trees[i].right != -1 {
		fmt.Println("right:", trees[i].right)
		setDepth(trees, depth, trees[i].right, d)
	}
	if trees[i].left != -1 {
		fmt.Println("left:", trees[i].left)
		// leftを深さ探索するとp(深さ)は+1
		// なぜ?
		setDepth(trees, depth, trees[i].left, d+1)
	}
}
