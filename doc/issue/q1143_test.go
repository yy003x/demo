package issue

import (
	"fmt"
	"testing"
)

/*
最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。
*/

func TestQ1143(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"
	rs := longestCommonSubsequence(text1, text2)
	fmt.Println(rs)
}

func longestCommonSubsequence(text1, text2 string) [][]int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for k := range dp {
		dp[k] = make([]int, n+1)
	}
	for i, r1 := range text1 {
		for j, r2 := range text2 {
			if r1 == r2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp
}
