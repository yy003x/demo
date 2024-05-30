package demo

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var tenToAny map[int]string = map[int]string{
	0: "z", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7",
	8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f",
	16: "g", 17: "h", 18: "j", 19: "k", 20: "m", 21: "n", 22: "p", 23: "q",
	24: "r", 25: "s", 26: "t", 27: "u", 28: "v", 29: "w", 30: "x", 31: "y",
}

func matrix() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover exit: ", err)
		}
	}()
	//获取当前时间
	t := time.Now()
	fmt.Println("app start:", t)

	//34636833
	fmt.Println(decimalToAny(33554432, 32))
	fmt.Println(anyToDecimal("1zzzzz", 32))
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
}

func findkey(in string) int {
	result := -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

func decimalToAny(num, n int) string {
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		if 32 > remainder && remainder > 9 {
			remainder_string = tenToAny[remainder]
		} else {
			remainder_string = strconv.Itoa(remainder)
		}
		new_num_str = remainder_string + new_num_str
		num = num / n
	}
	return new_num_str
}

func anyToDecimal(num string, n int) int {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findkey(value))
		if tmp != -1 {
			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(new_num)
}
