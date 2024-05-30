package demo

import (
	"fmt"
	"time"
)

func foo() {
	//defer cost(time.Now())

	//fmt.Println(accumulator(45))

	//go f1()
	go f2()
	ch := make(chan int)
	data, ok := <-ch
	fmt.Println(data)
	fmt.Println(ok)

	//close(ch)
	/*
		ch := make(chan int)
		for i := 48; i >= 0; i-- {
			go func(i int) {
				ch <- i
			}()
		}
		for data := range ch {
			fmt.Println(data)
			if data == 0 {
				break
			}
		}
	*/
}

func f1() {
	for {
		fmt.Println("call f1...")
	}
}

func f2() {
	fmt.Println("call f2...")
}

func cost(start time.Time) {
	tc := time.Since(start)
	fmt.Println(tc)
}

func accumulator(n int) (res int64) {
	switch n {
	case 1:
		res = 1
	case 2:
		res = 2
	default:
		res = accumulator(n-1) + accumulator(n-2)
	}
	return res
}
