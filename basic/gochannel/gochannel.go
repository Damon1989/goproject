package gochannel

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func ChannelOperate() {
	ch := make(chan int) //创建一个无缓冲通道
	// send
	go func() {
		ch <- 42 //向通道中发送数据
	}()
	// receive
	go func() {
		v := <-ch //从通道中接收数据
		println("Received:", v)
	}()

	time.Sleep(1 * time.Second)

	close(ch)
}

func ChannelFor() {
	/*ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

	// 使用for range遍历通道
	for v := range ch {
		println("Received:", v)
	}*/

	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(10)
		}
		close(ch)
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		for v := range ch {
			println("Received:", v)
		}
	}()

	wg.Wait()
}

func ChannelSync() {
	//初始化内容
	ch := make(chan int)
	fmt.Println("Length of channel:", len(ch), " Capacity of channel:", cap(ch))
	wg := sync.WaitGroup{}
	//间隔发送
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("Sending:", i, " at ", time.Now().Format("15:04:05.99999999"), "Length of channel:", len(ch), " Capacity of channel:", cap(ch))
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()
	//间隔接收
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("Received:", v, " at ", time.Now().Format("15:04:05.99999999"), "Length of channel:", len(ch), " Capacity of channel:", cap(ch))
			time.Sleep(3 * time.Second)
		}
	}()

	wg.Wait()
}

func ChannelASync() {
	//初始化内容
	ch := make(chan int, 5)
	fmt.Println("Length of channel:", len(ch), " Capacity of channel:", cap(ch))
	wg := sync.WaitGroup{}
	//间隔发送
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("Sending:", i, " at ", time.Now().Format("15:04:05.99999999"), "Length of channel:", len(ch), " Capacity of channel:", cap(ch))
			time.Sleep(1 * time.Second)
		}
		close(ch)
	}()
	//间隔接收
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("Received:", v, " at ", time.Now().Format("15:04:05.99999999"), "Length of channel:", len(ch), " Capacity of channel:", cap(ch))
			time.Sleep(3 * time.Second)
		}
	}()

	wg.Wait()
}

func ChannelGoroutineNumCtl() {
	// 1 独立的goroutine输出goroutine数量
	go func() {
		for {
			fmt.Println("NumGoroutine:", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 2 创建一个带缓冲区的通道，作为goroutine数量的控制器
	size := 1024
	ch := make(chan struct{}, size)

	for {
		// 一，启动goroutine前，执行 ch send 操作，占用一个缓冲区空间
		ch <- struct{}{}
		go func() {
			time.Sleep(10 * time.Second)
			// 二，goroutine结束前，执行 ch receive 操作，释放一个缓冲区空间
			<-ch
		}()
	}

}

func ChannelDirectional() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go getElement(ch, &wg)
	go setElement(ch, &wg, 100)

	wg.Wait()
}

func getElement(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("received from ch ,element is :", <-ch)
}

func setElement(ch chan<- int, wg *sync.WaitGroup, element int) {
	defer wg.Done()

	ch <- element
	fmt.Println("sent to ch ,element is :", element)
}

func ChannelForRange() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	for v := range ch {
		fmt.Println("Received:", v)
	}
}

func fn1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("sent to ch ,element is :", i)
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func fn2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("Received:", v)
	}
}

func ChannelReadWrite() {
	wg := sync.WaitGroup{}
	var ch = make(chan int, 10)
	wg.Add(2)
	go fn1(ch, &wg)
	go fn2(ch, &wg)
	wg.Wait()
}

func putNum(intChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	close(intChan)
	fmt.Println("close intChan")
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool, wg *sync.WaitGroup) {
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
	exitChan <- true
}

func printPrime(primeChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range primeChan {
		fmt.Println(v)
	}
}

func PrimeNumberChannel() {
	var wg sync.WaitGroup
	start := time.Now().UnixMilli()

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16)

	wg.Add(1)
	go putNum(intChan, &wg)

	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan, &wg)
	}

	wg.Add(1)
	go printPrime(primeChan, &wg)

	wg.Add(1)
	// 判断exitChan 是否存满值
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		close(primeChan)
		fmt.Println("close primeChan")
		wg.Done()
	}()
	fmt.Println("waiting for exitChan1")
	wg.Wait()
	fmt.Println("waiting for exitChan2")
	fmt.Println("执行完毕")
	end := time.Now().UnixMilli()
	fmt.Println("cost", end-start)
}
