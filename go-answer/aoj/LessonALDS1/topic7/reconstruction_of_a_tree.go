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

func searchIndex(array []int, target int) int {
	var ans int
	for i, a := range array {
		if a == target {
			ans = i
		}
	}
	return ans
}

func reconstruct(preorder, inorder []int) string {
	if len(preorder) == 0 {
		return ""
	}
	mid := preorder[0]
	splitPlace := searchIndex(inorder, mid)

	leftInorder := inorder[:splitPlace]
	leftPreorder := preorder[1 : len(leftInorder)+1]
	leftResult := reconstruct(leftPreorder, leftInorder)

	var result []string
	if leftResult != "" {
		result = append(result, leftResult)
	}

	rightInorder := inorder[splitPlace+1:]
	rightPreorder := preorder[len(leftInorder)+1:]
	rightResult := reconstruct(rightPreorder, rightInorder)

	if rightResult != "" {
		result = append(result, rightResult)
	}

	result = append(result, strconv.Itoa(mid))

	return strings.Join(result, " ")
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

	var preorder, inorder []int
	preorder = getIntList(scanner)
	inorder = getIntList(scanner)

	fmt.Println(reconstruct(preorder, inorder))
	writer.Flush()
}
