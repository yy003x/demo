package demo

import (
	"fmt"
	"strconv"
)

const ()

var (
	words = []string{"", "个", "十", "百", "千", "万", "十", "百", "千", "亿"}
	maps  = map[string]string{"0": "零", "1": "一", "2": "二", "3": "三", "4": "四", "5": "五", "6": "六", "7": "七", "8": "八", "9": "九"}
)

func Num2Word(num int) string {
	var ret string
	str := strconv.Itoa(num)
	len := len(str)
	if len > 9 {
		return "beyond limits"
	}
	for k, v := range str {
		nk := len - k
		fmt.Println(k)
		nv := string(v)
		// fmt.Println(string(v))
		if nv == "0" {
			ret += maps[string(nv)]
		} else {
			ret += maps[string(nv)] + words[nk]
		}
	}
	fmt.Println(num)
	fmt.Println(ret)
	return ret
}
func Th2Word(str string) string {
	return str
}
