package LessonALDS1

import "fmt"

func MergeSort(s []int) []int {
	var result []int
	// スライスの要素数が1の時return
	// returnされるとr := MergeSort(s[mid:])に行ける
	if len(s) < 2 {
		return s
	}

	mid := int(len(s) / 2)
	l := MergeSort(s[:mid])
	// 左側が全て要素数1になった時
	// 最後の要素でreturnになるとr := MergeSort(s[mid:])から開始し
	// 初期スライスの右側から再帰がスタートする
	r := MergeSort(s[mid:])
	i, j := 0, 0

	for i < len(r) && j < len(l) {
		if r[i] > l[j] {
			result = append(result, l[j])
			j++
		} else {
			result = append(result, r[i])
			i++
		}
	}

	result = append(result, r[i:]...)
	result = append(result, l[j:]...)

	return result
}

func main() {
	// Your code here!
	fmt.Println(MergeSort([]int{2, 1, 6, 5, 3, 8, 1}))
}
