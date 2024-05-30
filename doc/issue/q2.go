package issue

import "fmt"

/* @Title 两数之和

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

输入：nums = [2,7,11,15], target = 9
输出：[0,1]

*/
func Q2() {
	var arr = []int{3, 5, 7, 11, 13, 17, 19}
	var sum = 18
	rs := Q2Elegantly(arr, sum)
	fmt.Println(rs)

}

func Q2Elegantly(nums []int, sum int) (res []int) {
	var m = make(map[int]int)
	for k, v := range nums {
		if j, ok := m[sum-v]; ok {
			res = append(res, j, k)
			return
		}
		m[v] = k
	}
	return
}
func Q2Violently(nums []int, sum int) (res []int) {
	for k, v := range nums {
		for j := k + 1; j < len(nums); j++ {
			if v+nums[j] == sum {
				res = append(res, k, j)
				return
			}
		}
	}
	return
}
