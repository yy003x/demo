package base

import (
	"fmt"
	"testing"
)

type OrderSet struct {
	Elements []any
}

func NewOrderSet() *OrderSet {
	return &OrderSet{}
}

// Add 将元素添加到有序集合，保持升序排序
// Contains 检查元素是否在有序集合中
// Display 显示有序集合中的元素
// Remove 从集合中删除指定元素
// Size 获取集合的大小或长度
// Get 根据索引获取集合中的元素
// Clear 清空集合，删除所有元素

func StraightInsertionSort(slice []int) {
	n := len(slice)
	if n <= 1 {
		return // 切片已排序或为空，无需排序
	}

	for i := 1; i < n; i++ {
		key := slice[i] // 待插入的元素
		j := i - 1

		// 将大于 key 的元素右移
		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}

		// 插入 key 到正确的位置
		slice[j+1] = key
	}
}

// QuickSort 使用快速排序对切片进行排序
func QuickSort(slice []int) {
	if len(slice) <= 1 {
		return // 切片已排序或为空，无需排序
	}

	pivot := slice[0] // 选择第一个元素作为基准元素
	left, right := 1, len(slice)-1

	for left <= right {
		// 从左侧找到大于基准的元素
		for left <= right && slice[left] <= pivot {
			left++
		}

		// 从右侧找到小于基准的元素
		for left <= right && slice[right] >= pivot {
			right--
		}

		// 如果左侧元素大于右侧元素，交换它们
		if left <= right {
			slice[left], slice[right] = slice[right], slice[left]
		}
	}

	// 将基准元素放入正确的位置
	slice[0], slice[right] = slice[right], slice[0]

	// 递归排序基准元素左侧和右侧的子数组
	QuickSort(slice[:right])
	QuickSort(slice[right+1:])
}

func TestXxx(t *testing.T) {

	// 示例数据
	data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	fmt.Println("", data)
	QuickSort(data)
	fmt.Println("", data)

	// 直接插入排序
	// fmt.Println("", data)
	// StraightInsertionSort(data)
	// fmt.Println("", data)

}
