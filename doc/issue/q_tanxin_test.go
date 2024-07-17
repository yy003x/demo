package issue

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

// 买卖股票的最佳时机
func maxProfit(prices []int) int {
	n := len(prices)
	minPrice := 0
	maxProfit := 0
	for i := 0; i < n; i++ {
		if prices[i] <= minPrice || minPrice == 0 {
			minPrice = prices[i]
		} else {
			if prices[i]-minPrice > maxProfit {
				maxProfit = prices[i] - minPrice
			}
		}
	}
	return maxProfit
}

func TestCanJump(t *testing.T) {
	// nums := []int{2, 3, 1, 1, 4}
	// nums := []int{3, 2, 1, 0, 4}
	nums := []int{3, 2, 3, 1, 4, 0, 0, 1, 3, 2, 3, 0, 1}
	fmt.Println(canJump(nums))
}

// 跳跃游戏
func canJump(nums []int) bool {
	n := len(nums)
	step := 0
	for step < n && nums[step] > 0 {
		step = nums[step] + step
	}
	return step >= n-1
}

func TestJump(t *testing.T) {
	// nums := []int{2, 3, 1, 1, 4}
	// nums := []int{3, 2, 1, 0, 4}
	nums := []int{3, 2, 3, 1, 4, 0, 0, 1, 3, 2, 3, 0, 1}
	fmt.Println(jump(nums))
}

// 跳跃游戏2
func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
