package case_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestCh(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
}

// 定义一个对象池，用于管理*int类型的对象
var intPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("Allocating new object")
		return new(int)
	},
}

func TestPool(t *testing.T) {
	// 从池中获取一个对象
	obj := intPool.Get().(*int)
	fmt.Println("Got object:", *obj)
	// 使用对象
	*obj = 42
	fmt.Println("Updated object:", *obj)
	// 将对象放回池中
	intPool.Put(obj)
	// 再次从池中获取对象
	reusedObj := intPool.Get().(*int)
	fmt.Println("Got reused object:", *reusedObj)
	// 对象已经被重用，因此值为42
	*reusedObj = 100
	fmt.Println("Updated reused object:", *reusedObj)
	// 将对象放回池中
	intPool.Put(reusedObj)
}
