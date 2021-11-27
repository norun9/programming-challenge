package topic8

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
	return strings.Split(getNextLine(scanner), " ")
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

type Element struct {
	parent *Element
	left   *Element
	right  *Element
	key    int
}

type SearchBinaryTree struct {
	root *Element
}

func (t *SearchBinaryTree) insert(key int) {
	element := &Element{
		left:  nil,
		right: nil,
		key:   key,
	}

	var y *Element
	x := t.root
	//fmt.Println("現在のキー:", key)
	// 調査対象を深めていく
	for x != nil {
		y = x
		// 今回入ってきたkeyと調査対象のkeyを比較して
		// 今回のkeyの方が小さければ調査対象にxの左部分木を更新(調査を深める)
		if element.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
		//fmt.Println("今の親:", y.key)
	}
	if y != nil {
		//fmt.Println("最終の親:", y.key)
	}
	//fmt.Println("*******************")

	// 最終の親を現在のキーが入っているElementのparentフィールドに格納
	element.parent = y

	// y(親)がまだ無ければ今のElementを親(root)に設定する
	if y == nil {
		t.root = element
		// 親のkeyよりもElement(子)のkeyの数値が小さければ親の左部分木に設定
	} else if element.key < y.key {
		y.left = element
		// 親のkeyよりもElement(子)のkeyの数値が大きければ親の右部分木に設定
	} else {
		y.right = element
	}
}

func (t *SearchBinaryTree) find(k int, elem *Element) string {
	if elem == nil {
		return "no"
	}

	if elem.key == k {
		return "yes"
	}

	var exist string
	if k > elem.key {
		exist = t.find(k, elem.right)
	} else if k < elem.key {
		exist = t.find(k, elem.left)
	}

	return exist
}

func (t *SearchBinaryTree) preorder(elem *Element) {
	if elem == nil {
		return
	}
	fmt.Printf(" %d", elem.key)
	t.preorder(elem.left)
	t.preorder(elem.right)
}

func (t *SearchBinaryTree) inorder(elem *Element) {
	if elem == nil {
		return
	}
	t.inorder(elem.left)
	fmt.Printf(" %d", elem.key)
	t.inorder(elem.right)
}

func NewSearchBinaryTree() *SearchBinaryTree {
	return new(SearchBinaryTree)
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

	t := NewSearchBinaryTree()
	for i := 0; i < n; i++ {
		var A []string
		A = getStringList(scanner)
		if len(A) == 2 {
			cmd := A[0]
			key, _ := strconv.Atoi(A[1])

			switch cmd {
			case "insert":
				t.insert(key)
			case "find":
				fmt.Println(t.find(key, t.root))
			}
		} else {
			cmd := A[0]
			if cmd == "print" {
				t.inorder(t.root)
				fmt.Print("\n")
				t.preorder(t.root)
				fmt.Print("\n")
			}
		}
	}
	writer.Flush()
}
