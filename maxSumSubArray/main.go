package main

import (
	"fmt"
)

func main() {
	const MaxUint = ^uint(0)
	const MinUint = 0

	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	evenArr := []int{-2, 1, 3, -4, -5, 8, 3, 0, 6}
	oddArr := []int{-2, 1, 3, -4, 8, 3, 0, 6}
	l := len(oddArr)
	maxSum := MinInt
	if l == 0 {
		fmt.Println(0)
	} else if l == 1 {
		fmt.Println(1)
	} else {
		maxSum = max(maxSum, calcMaxSumSubArray(oddArr, 0, l-1, maxSum))
		fmt.Println(maxSum)
	}
}

func calcMaxSumSubArray(arr []int, startIdx int, endIdx int, maxSum int) int {
	mid := (startIdx + endIdx) / 2
	maxSum = max(maxSum, arr[mid]+arr[mid+1])
	maxSum = max(maxSum, arr[mid]+arr[mid-1])

}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
