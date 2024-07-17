package issue

import (
	"fmt"
	"testing"
)

/*
字母异位词分组

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

输入: strs = ["a"]
输出: [["a"]]
*/
func Test49(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ret := GroupAnagrams(strs)
	fmt.Println(ret)
}

func GroupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	hmap := make(map[int32][]string)
	for _, str := range strs {
		var long int32 = 0
		for _, rn := range str {
			long += rn
		}
		hmap[long] = append(hmap[long], str)
	}
	for _, v := range hmap {
		ret = append(ret, v)
	}
	return ret
}
