package issue

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	nums := []int{4, 10, 3, 5, 1, 2}
	// fmt.Println(InsertionSort(arr))
	// fmt.Println(MergeSort(arr))

	// fmt.Println(QuickSort(arr))
	fmt.Println(nums)
	fmt.Println(QuickSort(nums))
}

func HeapSort(nums []int) []int {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
	return nums
}

// 递归调整以索引 i 开始的子树为最大堆
func heapify(nums []int, n, i int) {
	largest := i
	left := i*2 + 1
	right := i*2 + 2
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	if right < n && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		heapify(nums, n, i)
	}
}

func QuickSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	pivot := partition(nums)
	QuickSort(nums[:pivot])
	QuickSort(nums[pivot+1:])
	return nums
}

func partition(nums []int) int {
	n := len(nums)
	pivot := n - 1
	start := 0
	for i := 0; i < n; i++ {
		if nums[i] < nums[pivot] {
			nums[start], nums[i] = nums[i], nums[start]
			start++
		}
	}
	nums[start], nums[pivot] = nums[pivot], nums[start]
	return start
}

func NewInsertionSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	mid := n / 2
	for i := mid; i < n; i++ {
		if arr[i] < arr[i-mid] {
			arr[i], arr[i-mid] = arr[i-mid], arr[i]
		}
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func SimpleQuickSort(arr []int) []int {
	ln := len(arr)
	if ln <= 1 {
		return arr
	}
	pivot := arr[0]
	left := []int{}
	right := []int{}
	for i := 1; i < ln; i++ {
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = SimpleQuickSort(left)
	right = SimpleQuickSort(right)
	// return append(append(QuickSort(left), pivot), QuickSort(right)...)
	return append(append(left, pivot), right...)
}

func MergeSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	mid := n / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:n])
	return merge(left, right)
}

func merge(left, right []int) []int {
	ln, rn := len(left), len(right)
	ans := make([]int, 0, ln+rn)
	l, r := 0, 0
	for l < ln && r < rn {
		if left[l] > right[r] {
			ans = append(ans, right[r])
			r++
		} else {
			ans = append(ans, left[l])
			l++
		}
	}
	ans = append(ans, left[l:]...)
	ans = append(ans, right[r:]...)
	return ans
}

func InsertionSort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		val := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > val {
			nums[j+1] = nums[j]
			fmt.Println(nums)
			j--
		}
		nums[j+1] = val
		fmt.Println(nums)
	}
	return nums
}

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
