package topic7

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
		nodes[i].left = 0
		nodes[i].right = -1
	}

	for iNode := 0; iNode < n; iNode++ {
		var A []int
		A = getIntList(scanner)
		id := A[0]

		for j, a := range A[1:] {
			if j == 0 {
				nodes[id].left = a
			} else {
				nodes[id].right = a
			}
			if a != -1 {
				nodes[a].parent = id
			}
		}
	}
	root := 0
	for nodes[root].parent != -1 {
		root = nodes[root].parent
	}

	//for _, node := range nodes {
	//	fmt.Println("parent:", node.parent)
	//	fmt.Println("left:", node.left)
	//	fmt.Println("right:", node.right)
	//	fmt.Println("***************")
	//}

	fmt.Print("Preorder\n")
	preOrder(nodes, root)
	fmt.Printf("\n")
	fmt.Print("Inorder\n")
	inOrder(nodes, root)
	fmt.Printf("\n")
	fmt.Print("Postorder\n")
	postOrder(nodes, root)
	fmt.Printf("\n")

	writer.Flush()
}

func inOrder(nodes []node, iNode int) {
	if iNode == -1 {
		return
	}
	inOrder(nodes, nodes[iNode].left)
	fmt.Printf(" %d", iNode)
	inOrder(nodes, nodes[iNode].right)
}

func preOrder(nodes []node, iNode int) {
	if iNode == -1 {
		return
	}
	node := nodes[iNode]
	fmt.Printf(" %d", iNode)
	preOrder(nodes, node.left)
	preOrder(nodes, node.right)
}

func postOrder(nodes []node, iNode int) {
	if iNode == -1 {
		return
	}
	leftNode := nodes[iNode].left
	postOrder(nodes, leftNode)
	rightNode := nodes[iNode].right
	postOrder(nodes, rightNode)
	fmt.Printf(" %d", iNode)
}
