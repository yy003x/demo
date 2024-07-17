package case_test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// a := []byte("AAAA/BBBBB")
	// fmt.Println(len(a))
	// fmt.Println(cap(a))
	// idx := bytes.IndexByte(a, '/')
	// b := a[:idx]
	// fmt.Println(len(b))
	// fmt.Println(cap(b))
	// c := a[idx+1:]
	// fmt.Println(len(c))
	// fmt.Println(cap(c))
	// // c = append(c, "DD"...)
	// b = append(b, "CCC"...)
	// fmt.Println(len(b))
	// fmt.Println(cap(b))
	// fmt.Println(string(a))
	// fmt.Println(string(b))
	// fmt.Println(string(c))
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	idx := 5
	b := nums[:idx]
	c := nums[idx+1:]
	b = append(b, []int{6, 6, 6}...)
	fmt.Println(nums)
	fmt.Println(b)
	fmt.Println(c)
}

func TestUint(t *testing.T) {
	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
}
