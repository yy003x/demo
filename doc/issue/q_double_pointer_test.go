package issue

import (
	"fmt"
	"sort"
	"testing"
)

// 颜色分类
func SortColors(nums []int) {
	n := len(nums)
	p0, p2 := 0, n-1
	for i := 0; i <= p2; i++ {
		for i <= p2 && nums[i] == 2 {
			nums[p2], nums[i] = nums[i], nums[p2]
			p2--
		}
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
		}
	}
}

// 三数之和
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := [][]int{}
	// 枚举a
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		// 枚举b c
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}

func TestPointer(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0, 2, 0, 0}
	SortColors(nums)
	fmt.Println(nums)
	fmt.Println(ThreeSum(nums))
}
