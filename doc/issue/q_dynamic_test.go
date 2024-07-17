package issue

import (
	"testing"
)

// 最长回文子串
func LongestPalindrome(s string) string {
	ln := len(s)
	if ln == 0 {
		return ""
	}
	expand := func(s string, left, right int) (int, int) {
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				left--
				right++
			} else {
				break
			}
		}
		return left + 1, right - 1
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

// 最长公共子序列
func LongestCommonSubsequence(text1, text2 string) [][]int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for k := range dp {
		dp[k] = make([]int, n+1)
	}
	for i, t1 := range text1 {
		for j, t2 := range text2 {
			if t1 == t2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp
}

/*
编辑距离
 v - 	r 	o 	s
 - k 	01	02 	03
 h 10	1  	2	3
 o 20	2  	1	2
 r 30	2	2	2
 s 40	3	3	2
 e 50	4	5	3
*/
func MinDistance(word1, word2 string) [][]int {
	m, n := len(word1), len(word2)
	ans := make([][]int, m+1)
	for line := range ans {
		ans[line] = make([]int, n+1)
		ans[line][0] = line
	}
	for col := 0; col <= n; col++ {
		ans[0][col] = col
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				ans[i][j] = ans[i-1][j-1]
			} else {
				ans[i][j] = min(ans[i-1][j-1], min(ans[i-1][j], ans[i][j-1])) + 1
			}
		}
	}
	return ans
}

// 不同路径
func UniquePaths(m, n int) [][]int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp
}

// 最小路径和
func MinPathSum(grid [][]int) [][]int {
	for i, line := range grid {
		for j := range line {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}
	return grid
}

// 杨辉三角
func GenerateYangHui(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}

// 斐波那契数列
func Fibonacci(n int) []int {
	ret := []int{}
	if n <= 0 || n >= 45 {
		return ret
	}
	a := 0
	b := 1
	for i := 1; i <= n; i++ {
		a, b = b, a+b
		ret = append(ret, b)
	}
	return ret
}

func TestDynamic(t *testing.T) {
	// 斐波那契数列
	// fmt.Println(Fibonacci(8))

	// 杨辉三角
	// fmt.Println(GenerateYangHui(5))
	// 最长回文
	// fmt.Println(LongestPalindrome("xxxabcdeffedcbazzz"))

	// 最长子序列
	// fmt.Println(LongestCommonSubsequence("abcde", "ace"))

	//编辑距离
	// fmt.Println(MinDistance("horse", "ros"))
	// fmt.Println(MinDistance("intention", "execution"))

	// 不同路径
	// fmt.Println(UniquePaths(3, 7))
	// 最小路径和
	// fmt.Println(MinPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
