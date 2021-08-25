package main

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

type node struct {
	parent int
	lft    int
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

	nodes := make([]node, n)

	for i := 0; i < n; i++ {
		nodes[i].parent = -1
		nodes[i].left = -1
		nodes[i].right = -1
	}

	for i := 0; i < n; i++ {
		var C []int
		C = getIntList(scanner)
		id := C[0]
		for j, c := range C[1:] {
			if j == 0 {
				nodes[id].left = c
			} else {
				nodes[id].right = c
			}
			if c != -1 {
				nodes[c].parent = id
			}
		}
	}

	depth := make([]int, n)
	for i, node := range nodes {
		if node.parent == -1 {
			setDepth(nodes, depth, i, 0)
			break
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("node %d: parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s\n", i, nodes[i].parent, setSibling(nodes, i), getDegree(nodes[i]), depth[i], setHeight(nodes, i), getType(nodes[i]))
	}

	writer.Flush()
}

func setHeight(nodes []node, i int) int {
	if i == -1 {
		return -1
	}
	leftID := nodes[i].left
	leftHeight := setHeight(nodes, leftID)

	rightID := nodes[i].right
	rightHeight := setHeight(nodes, rightID)

	return max(leftHeight, rightHeight) + 1
}

func setSibling(nodes []node, i int) int {
	parentID := nodes[i].parent
	if parentID == -1 {
		return -1
	}

	parentNode := nodes[parentID]
	if parentNode.left == i {
		return parentNode.right
	}
	return parentNode.left
}

func getDegree(n node) int {
	var degree int
	if n.left != -1 {
		degree++
	}
	if n.right != -1 {
		degree++
	}
	return degree
}

func getType(n node) string {
	if n.parent == -1 {
		return "root"
	}
	if n.left == -1 && n.right == -1 {
		return "leaf"
	}

	return "internal node"
}

func setDepth(nodes []node, depth []int, i, p int) {
	depth[i] = p

	if nodes[i].right != -1 {
		setDepth(nodes, depth, nodes[i].right, p+1)
	}
	if nodes[i].left != -1 {
		setDepth(nodes, depth, nodes[i].left, p+1)
	}
}
