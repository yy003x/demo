package issue

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(FindDuplicate(nums))
}

// 寻找重复数 快慢指针
func FindDuplicate(nums []int) int {
	slow := 0
	fast := 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// 只出现一次的数字
// a^b^a=a^a^b=0^b=b
func SingleNumber(nums []int) int {
	single := 0
	for _, v := range nums {
		single ^= v
	}
	return single
}

// 多数元素
// hash
func MajorElement(nums []int) int {
	hmap := make(map[int]int)
	for _, v := range nums {
		hmap[v]++
		if hmap[v] > len(nums)/2 {
			return v
		}
	}
	return -1
}

// 摩尔投票
func MooreElement(nums []int) int {
	count := 0
	candidate := 0
	//投票
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if candidate == v {
			count++
		} else {
			count--
		}
	}
	//验证
	count = 0
	for _, v := range nums {
		if candidate == v {
			count++
		}
	}
	if count > len(nums)/2 {
		return candidate
	}
	return -1
}
