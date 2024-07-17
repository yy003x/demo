package issue

import (
	"fmt"
	"testing"
)

/*
多数元素
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

输入：nums = [3,2,3]
输出：3
输入：nums = [2,2,1,1,1,2,2]
输出：2
*/

func Test169(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
	fmt.Println(mooreElement(nums))
}

func majorityElement(nums []int) int {
	hmap := make(map[int]int)
	for _, num := range nums {
		hmap[num]++
	}
	for k, v := range hmap {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

func mooreElement(nums []int) int {
	count := 0
	vote := 0
	for _, num := range nums {
		if count == 0 {
			vote = num
		}
		if vote == num {
			count++
		} else {
			count--
		}
	}
	return vote
}
