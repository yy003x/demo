package issue

import (
	"fmt"
	"testing"
)

/*
最长连续序列

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
*/
func Test128(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2, 3, 4, 6, 5, 4, 101, 100}
	result := longestConsecutive(nums)
	fmt.Printf("最长连续序列的长度为：%d\n", result)
}

func longestConsecutive(nums []int) int {
	hmap := make(map[int]bool)
	for _, v := range nums {
		hmap[v] = true
	}
	longestLen := 0
	fmt.Println(hmap)
	for num := range hmap {
		if !hmap[num-1] {
			currentNum := num
			currentLen := 1
			for hmap[currentNum+1] {
				currentNum++
				currentLen++
			}
			longestLen = max(longestLen, currentLen)
		}
	}
	return longestLen
}
