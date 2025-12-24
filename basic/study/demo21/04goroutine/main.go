package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func test(num int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("协程(%v)打印的第%v条数据\n", num, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
}
