package case_test

import (
	"fmt"
	"testing"
)

func TestClosure(tt *testing.T) {
	fmt.Println("func1 rtn", func1())
	fmt.Println("func2 rtn", func2())
	fmt.Println("func3 rtn", func3())
	fmt.Println("func4 rtn", func4())
	fmt.Println("func5 rtn", func5())
}

func func1() (i int) {
	i = 1
	defer func() {
		i++
		fmt.Println("func1 defer1", i)
	}()
	defer func() {
		i++
		fmt.Println("func1 defer2", i)
	}()
	i++
	return
}

func func2() int {
	i := 1
	defer func() {
		i++
		fmt.Println("func2 defer1", i)
	}()
	defer func() {
		i++
		fmt.Println("func2 defer2", i)
	}()
	i++
	return i
}
func func3() int {
	i := 1
	defer func(i int) {
		i++
		fmt.Println("func3 defer1", i)
	}(i)
	defer func(i int) {
		i++
		fmt.Println("func3 defer2", i)
	}(i)
	i++
	return i
}

func func4() (i int) {
	i = 1
	defer func(i int) {
		i++
		fmt.Println("func4 defer1", i)
	}(i)
	defer func(i int) {
		i++
		fmt.Println("func4 defer2", i)
	}(i)
	i++
	return
}

func func5() (i int) {
	i = 1
	defer func() {
		i := func() int {
			return i * 2
		}()
		fmt.Println("func5 defer", i)
	}()
	i++
	return
}
