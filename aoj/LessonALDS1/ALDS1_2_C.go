package LessonALDS1

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func insertionSort2(arrOrig []int, N int) []int {
	A := make([]int, N)
	copy(A, arrOrig)
	for i := 1; i < N; i++ {
		x := A[i]
		j := i - 1
		for ; j >= 0 && A[j] > x; j-- {
			// 値を後ろにずらしていく
			A[j+1] = A[j]
		}
		// 重複していても最後に重複しあたところに値が入るから大丈夫
		A[j+1] = x
		fmt.Println(i, "finit")
		fmt.Println(A)
		fmt.Println("===")
	}
	return A
}

func insertionSort(arrOrig []string, N int) []string {
	n := len(arrOrig)

	A := make([]string, n)
	copy(A, arrOrig)
	for i := 1; i < n; i++ {
		v := A[i]
		j := i - 1
		for j >= 0 && A[j][1] > v[1] {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = v
	}
	return A
}

func bubbleSort(arrOrig []string, N int) []string {
	n := len(arrOrig)

	A := make([]string, n)
	copy(A, arrOrig)
	flag := true
	for flag {
		flag = false
		for i := N - 1; i >= 1; i-- {
			num1, _ := strconv.Atoi(strings.Split(A[i], "")[1])
			num2, _ := strconv.Atoi(strings.Split(A[i-1], "")[1])
			if num1 < num2 {
				flag = true
				A[i], A[i-1] = A[i-1], A[i]
			}
		}
	}
	return A
}

func selectionSort(arrOrig []string, N int) []string {
	n := len(arrOrig)

	A := make([]string, n)
	copy(A, arrOrig)
	for i := 0; i < N; i++ {
		min := i
		for j := i; j < N; j++ {
			num1, _ := strconv.Atoi(strings.Split(A[min], "")[1])
			num2, _ := strconv.Atoi(strings.Split(A[j], "")[1])
			if num1 > num2 {
				min = j
			}
		}
		A[i], A[min] = A[min], A[i]
	}
	return A
}

func printStableOrNot(arr1, arr2 []string) {
	if reflect.DeepEqual(arr1, arr2) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	input := strings.Split(sc.Text(), " ")
	base := insertionSort(input, n)

	result1 := bubbleSort(input, n)
	fmt.Println(strings.Trim(fmt.Sprintf("%v", result1), "[]"))
	printStableOrNot(base, result1)
	result2 := selectionSort(input, n)
	fmt.Println(strings.Trim(fmt.Sprintf("%v", result2), "[]"))
	printStableOrNot(base, result2)
}
