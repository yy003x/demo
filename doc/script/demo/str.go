package demo

import (
	"be_demo/doc/script/base"
	"fmt"
	"strings"
)

//利用 栈 DefStack 去除相邻两个重复的字符串
func RemoveDuplicates(s string) {
	st := base.NewDefStack()
	for _, v := range s {
		if st.Top != nil && st.Top.Value == string(v) {
			st.Pop()
		} else {
			st.Push(string(v))
		}
	}
	fmt.Println(s)
	st.Show()
}

//简单 + 实用 去除相邻两个重复的字符串
func RemoveRepeat(s string) string {
	arr := strings.Split(s, "")
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			arr = append(arr[:i-1], arr[i+1:]...)
			if i < 2 {
				i = 0
			} else {
				i -= 2
			}
		}
	}
	ret := strings.Join(arr, "")
	fmt.Println(s)
	fmt.Println(ret)
	return ret
}
