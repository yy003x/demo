package issue

import (
	"fmt"
	"testing"
)

/*
两数之和

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
*/
func TestQ1(t *testing.T) {
	var arr = []int{3, 5, 7, 11, 13, 17, 19}
	var sum = 18
	rs := TwoSum(arr, sum)
	fmt.Println(rs)
}

func TwoSum(arr []int, sum int) []int {
	hmap := make(map[int]int)
	for key, val := range arr {
		if one, ok := hmap[sum-val]; ok {
			return []int{one, key}
		}
		hmap[val] = key
	}
	return []int{}
}

func Q2Violently(arr []int, sum int) []int {
	for k, v := range arr {
		for j := k + 1; j < len(arr); j++ {
			if v+arr[j] == sum {
				return []int{v, arr[j]}
			}
		}
	}
	return []int{}
}
