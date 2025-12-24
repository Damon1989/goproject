package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

// 写的方法
func write() {
	mutex.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}

// 读的方法
func read() {
	mutex.RLock()
	fmt.Println("---执行读操作")
	time.Sleep(time.Second * 2)
	mutex.RUnlock()
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
}
