package base

import (
	"fmt"
	"testing"
)

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func SelectSort(arr []int) []int {
	return arr
}

func TestSort(t *testing.T) {
	arr := []int{90, 34, 25, 12, 22, 64, 11}
	fmt.Println(arr)
	// BubbleSort(arr)
	fmt.Println(arr)
}
