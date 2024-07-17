package issue

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestDailyTemperaturesx(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}

// 每日温度
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			day := stack[len(stack)-1]
			ans[day] = i - day
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func TestDecodeString(t *testing.T) {
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
}

//
func decodeString(s string) string {
	stack := []string{}
	num := 0
	currStr := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= '0' && char <= '9' {
			num = int(char - '0')
			fmt.Println(num)
		} else if char == '[' {
			stack = append(stack, currStr)
			stack = append(stack, strconv.Itoa(num))
			currStr = ""
		} else if char == ']' {
			count, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			prevStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			currStr = prevStr + strings.Repeat(currStr, count)
		} else {
			currStr += string(char)
		}
	}
	return currStr
}
