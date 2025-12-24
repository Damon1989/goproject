package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	defer wg.Done()
	for i := 2; i <= 120000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	defer wg.Done()
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	//close(primeChan)  // 如果一个channel关闭了就没法给这个channel发送数据了
	// 给exitChan里面放入一条数据
	exitChan <- true
}

// printPrime打印素数的方法
func printPrime(primeChan chan int) {
	defer wg.Done()
	for {
		num, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(num)
	}
}

func main() {
	start := time.Now().UnixMicro()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16) // 标识primeChan close

	// 存放素数的协程
	wg.Add(1)
	go putNum(intChan)

	// 统计素数的协程
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	// 打印素数的协程
	wg.Add(1)
	go printPrime(primeChan)

	// 判断exitChan是否存满值

	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	end := time.Now().UnixMicro()
	fmt.Println("执行完毕...", end-start)
}
