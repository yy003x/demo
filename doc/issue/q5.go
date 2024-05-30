package issue

import "fmt"

/* @Title

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
*/

func Q5() {
	var arr = []int{2, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(Q5Elegantly(arr))
}
func TrapRain(arr []int) (rs int) {
	return
}
func Q5Elegantly(arr []int) (rs int) {
	var size = len(arr)
	if size == 0 {
		return
	}
	var maxLeft = make([]int, size)
	var maxRight = make([]int, size)
	maxLeft[0] = arr[0]
	for i := 1; i < size; i++ {
		maxLeft[i] = max(arr[i], maxLeft[i-1])
	}
	maxRight[size-1] = arr[size-1]
	for j := size - 2; j >= 0; j-- {
		maxRight[j] = max(arr[j], maxRight[j+1])
	}
	fmt.Println(maxLeft, maxRight)
	for k := 1; k < size-1; k++ {
		rs += min(maxLeft[k], maxRight[k]) - arr[k]
	}
	return
}

func Q5Violently(arr []int) (rs int) {
	var size = len(arr)
	if size == 0 {
		return
	}
	for i := 1; i < size-1; i++ {
		var maxLeft, maxRight int
		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, arr[j])
		}
		for k := i; k < size; k++ {
			maxRight = max(maxRight, arr[k])
		}
		rs += min(maxLeft, maxRight) - arr[i]
	}
	return
}
