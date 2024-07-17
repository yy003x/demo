package issue

import (
	"fmt"
	"testing"
)

// 盛最多水的容器
func MaxArea(height []int) int {
	n := len(height)
	left, right := 0, n-1
	ans := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		ans = max(ans, area)
	}
	return ans
}

// 接雨水
func TrapRain(height []int) int {
	n := len(height)
	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	ans := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}
func TestWater(t *testing.T) {
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// fmt.Println(MaxArea(height))
	nums1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	nums2 := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(TrapRain(nums1))
	fmt.Println(TrapRain(nums2))
}
