package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Tennis struct {
	HitChan chan int
	Wg      sync.WaitGroup
	Rander  *rand.Rand
}

func NewTennisGame() *Tennis {
	return &Tennis{
		HitChan: make(chan int),
		Rander:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (t *Tennis) Player(name string) {
	t.Wg.Add(1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("unknown panic info: %v \n", e)
			}
			t.Wg.Done()
		}()
		for {
			// time.Sleep(time.Second)
			ball, ok := <-t.HitChan
			if !ok {
				fmt.Printf("Player %s Won\n", name)
				return
			}
			n := t.Rander.Intn(20)
			fmt.Println(n)
			if n%5 == 0 {
				fmt.Printf("Player %s Missed\n", name)
				close(t.HitChan)
				return
			}
			fmt.Printf("Player %s Hit %d\n", name, ball)
			ball++
			t.HitChan <- ball
		}
	}()
}
func (t *Tennis) Begin() {
	t.HitChan <- 1
	t.Wg.Wait()
}
