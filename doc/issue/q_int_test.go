package issue

import (
	"fmt"
	"math"
	"testing"
)

/*
 整数反转
 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

输入：x = 123
输出：321

输入：x = -123
输出：-321

输入：x = 120
输出：21
*/

func TestInt(t *testing.T) {
	// fmt.Println(ReverseInt(123))
	// fmt.Println(ReverseInt(-123))
	// fmt.Println(ReverseInt(120))
	// fmt.Println(ReverseStr("hello world!"))
	nums := []int{1, 3, 8, 6, 5, 4, 2}
	fmt.Println(nums)
	NextPermutation(nums)
	fmt.Println(nums)
}

func ReverseStr(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}

func NextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// Step 1: Find the first decreasing element from the end
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	fmt.Println(i)
	// If such an element was found
	if i >= 0 {
		// Step 2: Find the element just larger than nums[i] from the end
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// Step 3: Swap elements at i and j
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println(i)
	fmt.Println(nums)
	// Step 4: Reverse the elements after index i
	reverse(nums[i+1:])
}

// reverse reverses the slice in place.
func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func ReverseInt(x int) int {
	ret := 0
	for x != 0 {
		if x > math.MaxInt32 || x < math.MinInt32 {
			return 0
		}
		i := x % 10
		ret = i + ret*10
		x = x / 10
	}
	return ret
}
