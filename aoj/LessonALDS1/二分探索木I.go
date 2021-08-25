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

type BinarySearchTree struct {
	root *Element
}

func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

func (t *BinarySearchTree) insert(key int) {
	element := &Element{
		left:  nil,
		right: nil,
		key:   key,
	}

	var y *Element
	x := t.root
	fmt.Println("現在のキー:", key)
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
		fmt.Println("今の親:", y.key)
	}
	if y != nil {
		fmt.Println("最終の親:", y.key)
	}
	fmt.Println("*******************")

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

func (t *BinarySearchTree) preorderTreeWalk(p *Element) {
	if p == nil {
		return
	}
	fmt.Printf(" %d", p.key)
	t.preorderTreeWalk(p.left)
	t.preorderTreeWalk(p.right)
}

func (t *BinarySearchTree) inorderTreeWalk(p *Element) {
	if p == nil {
		return
	}

	t.inorderTreeWalk(p.left)
	fmt.Printf(" %d", p.key)
	t.inorderTreeWalk(p.right)
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
	writer.Flush()

	bstree := NewBinarySearchTree()
	for i := 0; i < n; i++ {
		var A []string
		A = getStringList(scanner)
		cmd := A[0]
		if cmd == "insert" {
			k, _ := strconv.Atoi(A[1])
			bstree.insert(k)
		}
		if cmd == "print" {
			bstree.inorderTreeWalk(bstree.root)
			fmt.Print("\n")
			bstree.preorderTreeWalk(bstree.root)
			fmt.Print("\n")
		}
	}

	// inorderが呼ばれる順番
	// 再帰関数
	//fmt.Println(bstree.root.left.left.key)
	//fmt.Println(bstree.root.left.key)
	//fmt.Println(bstree.root.left.right.left.key)
	//fmt.Println(bstree.root.left.right.key)
	//fmt.Println(bstree.root.left.right.right.key)
	//fmt.Println(bstree.root.key)
	//fmt.Println(bstree.root.right.key)
}
