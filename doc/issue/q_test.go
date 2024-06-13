package issue

import (
	"fmt"
	"testing"
)

func TestQ(t *testing.T) {
	single := 0
	single = single ^ 3
	single ^= 0
	fmt.Println(single)
}

func QuckSort(arr []int) []int {
	ln := len(arr)
	if ln == 0 {
		return arr
	}
	return []int{}
}
