package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getScanner(fp *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 500001), 500000)
	return scanner
}

func getNextString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func getNextInt(scanner *bufio.Scanner) int {
	i, _ := strconv.Atoi(getNextString(scanner))
	return i
}

func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}

	scanner := getScanner(fp)
	writer := bufio.NewWriter(os.Stdout)

	n := getNextInt(scanner)

	trees := make([]Tree, n)
	depth := make([]int, n)
	for i := 0; i < n; i++ {
		trees[i].parent = -1
		trees[i].left = -1
		trees[i].right = -1
	}
	for i := 0; i < n; i++ {
		id := getNextInt(scanner)
		k := getNextInt(scanner)
		//fmt.Println("id:", id, "k:", k)
		l := -1
		for ii := 0; ii < k; ii++ {
			c := getNextInt(scanner)
			//fmt.Println("c:", c)
			trees[c].parent = id
			if ii == 0 {
				// 最初の子ノードはleft_childに格納
				trees[id].left = c
			} else {
				// 次からは最初の子ノードの兄弟なので、right_siblingに入れる。
				trees[l].right = c
			}
			// 親ノードの番号を入れる
			l = c
			//fmt.Println("l:", l)
		}
	}
	//for idx, t := range trees {
	//	fmt.Println("index:", idx)
	//	fmt.Println("parent:", t.parent)
	//	fmt.Println("left:", t.left)
	//	fmt.Println("right:", t.right)
	//	fmt.Println("********************")
	//}
	for i := 0; i < n; i++ {
		if trees[i].parent == -1 {
			fmt.Println("depth:", depth)
			setDepth(trees, depth, i, 0)
			break
		}
	}

	fmt.Println("depth:", depth)

	buf := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, fmt.Sprintf("node %d: parent = %d, depth = %d, %s, [%s]", i, trees[i].parent, depth[i], getType(trees[i]), getChildren(trees, i, buf)))
	}

	writer.Flush()
}

func setDepth(trees []Tree, depth []int, i, p int) {
	// 深さを設定する
	depth[i] = p
	if trees[i].right != -1 {
		fmt.Println("right:", trees[i])
		// 右の兄弟に同じ深さを設定 tree[i]は左の兄弟が来る為、深さはそのまま
		setDepth(trees, depth, trees[i].right, p)
	}
	if trees[i].left != -1 {
		fmt.Println("left:", trees[i].left)
		// leftを深さ探索するとp(深さ)は+1
		// 最も左の子に自分の深さ+1を設定
		setDepth(trees, depth, trees[i].left, p+1)
	}
}

func getType(tree Tree) string {
	if tree.parent == -1 {
		return "root"
	}
	if tree.left == -1 {
		return "leaf"
	}
	return "internal node"
}

func getChildren(trees []Tree, i int, buf []string) string {
	end := 0
	for c := trees[i].left; c != -1; c = trees[c].right {
		buf[end] = fmt.Sprintf("%d", c)
		end++
	}
	return strings.Join(buf[0:end], ", ")
}

type Tree struct {
	// 親のノード
	parent int
	// 子のノード
	left int
	// 左の兄弟のノード
	right int
}
