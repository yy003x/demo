package case_test

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"
)

func TestAC(t *testing.T) {
	m := map[string]float32{"hello": 1}
	fmt.Println(m["hello"])
	fmt.Println(m["world"])
}

func TestB(t *testing.T) {
	a := new([16]byte)
	for i := 1; i < 10; i++ {
		rand.Read(a[:])
		fmt.Println(a)
	}
}

func TestA(t *testing.T) {
	c := make(chan *[16]byte)
	go func() {
		// 使用两个数组以避免数据竞争。
		var dataA, dataB = new([16]byte), new([16]byte)
		for {
			_, err := rand.Read(dataA[:])
			if err != nil {
				close(c)
			} else {
				c <- dataA
				dataA, dataB = dataB, dataA
			}
		}
	}()
	for data := range c {
		fmt.Println((*data)[:])
		time.Sleep(time.Second / 2)
	}
}

func TestC(t *testing.T) {
	addone := func(x int) int { return x + 1 }
	square := func(x int) int { return x * x }
	double := func(x int) int { return x + x }
	transforms := map[string][]func(int) int{
		"inc,inc,inc": {addone, addone, addone},
		"sqr,inc,dbl": {square, addone, double},
		"dbl,sqr,sqr": {double, double, square},
	}
	for _, n := range []int{2, 3, 5, 7} {
		fmt.Println(">>>", n)
		for name, transfers := range transforms {
			result := n
			for _, xfer := range transfers {
				result = xfer(result)
			}
			fmt.Printf(" %v: %v \n", name, result)
		}
	}
}
