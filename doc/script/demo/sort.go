package demo

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateList(n int) (list []int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		list = append(list, rand.Intn(100))
	}
	return
}

// 数组中重复的数字
func Topic3() {
	// 交换函数
	swap := func(list []int, i, j int) {
		list[i], list[j] = list[j], list[i]
	}

	input := []int{2, 3, 1, 0, 2, 5}
	for i := 0; i < len(input); i++ {
		// 将转换后
		for i != input[i] {

			// i要换到input[i]的位置，但是如果input[i]位置的值（intput[input[i]]）等与input[i]，说明重复来了
			if input[i] == input[input[i]] {
				fmt.Println(input[i], "重复")
				goto TASK
			}
			swap(input, i, input[i])
		}
	}
TASK:
	fmt.Println("done")
}
