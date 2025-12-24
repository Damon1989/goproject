package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fn1(ch chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("【写入】数据：%v\n", i)
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}

func fn2(ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("【读取】数据：%v\n", v)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	var ch = make(chan int, 10)
	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)
	wg.Wait()
}
