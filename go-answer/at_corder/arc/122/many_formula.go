package _22

import (
	"bufio"
	"fmt"
	"math"
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

// AC 6/11
func main() {
	fp := os.Stdin
	if len(os.Args) > 1 {
		fp, _ = os.Open(os.Args[1])
	}
	scanner := bufio.NewReader(fp)
	writer := bufio.NewWriter(os.Stdout)
	var k int
	var A []int
	fmt.Sscan(getNextLine(scanner), &k)
	A = getIntList(scanner)
	base := (math.Pow(10, 9) + float64(7))

	if k == 1 {
		fmt.Println(A[0] % int(base))
		return
	}
	result := 0
	for bit := uint(0); bit < 1<<uint(k-1); bit++ {
		sum := A[0]
		//fmt.Println("**************************")
		//fmt.Printf("bit: %08b\n", bit)

		isValid := true
		var isDouble bool
		for i := uint(0); i < uint(k-1); i++ {
			//fmt.Printf("index: %08b\n", i)
			if bit&(1<<i) != 0 {
				sum += A[i+1]
				//fmt.Println("+", A[i+1])
				isDouble = false
			} else {
				if isDouble {
					isValid = false
					break
				}
				sum -= A[i+1]
				isDouble = true
			}
		}
		if !isValid {
			sum = 0
		}
		//fmt.Println(sum)
		result += sum
	}

	fmt.Printf("%d\n", result%int(base))
	writer.Flush()
}
