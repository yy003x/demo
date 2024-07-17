package case_test

import "testing"

type Node struct {
	Next *Node
}

var ch = make(chan int)

func goroutine1() {
	ch <- 1 // 阻塞，等待goroutine2接收
}

func goroutine2() {
	<-ch // 阻塞，等待goroutine1发送
}

func TestSort(t *testing.T) {
	go goroutine1()
	go goroutine2()

}
