package demo

import (
	//"os"
	"fmt"
	"strings"
)

func gbyte() {

	ten := 1048576
	fmt.Println(ten)
	hex := make([]int, 0)
	for {
		m := ten % 16
		ten = ten / 16
		fmt.Println(m)
		fmt.Println(ten)

		if ten == 0 {
			hex = append(hex, m)
			break
		}
		hex = append(hex, m)
	}
	fmt.Println(hex)

	hexArr := []string{}
	for i := len(hex) - 1; i >= 0; i-- {
		if hex[i] >= 10 {
			hexArr = append(hexArr, fmt.Sprintf("%c", 'a'+hex[i]-10))
		} else {
			hexArr = append(hexArr, fmt.Sprintf("%d", hex[i]))
		}
	}
	hexStr := strings.Join(hexArr, "")
	fmt.Println(hexStr)

	//b := a << 8
	//c := b >> 8
}
