package LessonALDS1

import (
	"fmt"
)

func shellSort(arrOrig []int, N int) {
	A := make([]int, N)
	copy(A, arrOrig)
	gap := N / 2
	for gap > 0 {
		// gapから配列の最後まで見ていく
		for i := gap; i < N; i++ {
			// 現在調べている要素をtmp保管する
			tmp := A[i]
			j := i
			// jがgap以上かつA[j-gap]がtmpよりも大きい値まで
			for j >= gap && A[j-gap] > tmp {
				A[j] = A[j-gap]
				j -= gap
			}
			// 最後に終了したgapの数値を添字にして
			A[j] = tmp
		}
		// 現在のgapを2で割って1まで減らしていく
		gap /= 2
	}
	fmt.Println(A)
}
