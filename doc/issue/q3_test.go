package issue

import (
	"fmt"
	"testing"
)

/*
无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func TestQ3(t *testing.T) {
	s := "pwwkewacbdefwkefa"
	fmt.Println(LongestSubStr(s))
}

func LongestSubStr(s string) int {
	var (
		left = 0
		long = 0
		hmap = make(map[byte]int)
	)
	for right := 0; right < len(s); right++ {
		if key, ok := hmap[s[right]]; ok {
			left = max(left, key+1)
		}
		long = max(long, right-left+1)
		hmap[s[right]] = right
	}
	return long
}

func Q1Violently(s string) int {
	var l = 0
	for i := 0; i < len(s); i++ {
		var rs = ""
		var m = make(map[string]int)
		for j := i; j < len(s); j++ {
			str := string(s[j])
			if _, ok := m[str]; ok {
				break
			}
			m[str] = j
			rs += str
		}
		fmt.Println(rs)
		l = max(l, len(rs))
	}
	return l
}
