package case_test

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"
)

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
