package issue

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	// rs := TwoSum([]int{1, 3, 4, 5, 6}, 12)
	arr := []int{123, 23, 4, 5, 234, 1, 234, 5, 234, 5, 21, 43, 4, 2142, 4, 42, 5, 23, 52, 31}
	fmt.Println(arr)
	// BubbleSort(arr)
	// SelectionSort(arr)
	InsertionSort(arr)
	fmt.Println(arr)
}

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = val
	}
}

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
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
