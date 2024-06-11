package issue

import (
	"fmt"
	"testing"
)

/* @Title 给你一个字符串 s，找到 s 中最长的 回文子串

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/

func TestQ5(t *testing.T) {
	s := "xxxabcdeffedcbazzz"
	fmt.Println(long(s))
}

func long(s string) string {
	ln := len(s)
	if ln == 0 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < ln; i++ {
		left1, right1 := expand(s, i, i)
		left2, right2 := expand(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
