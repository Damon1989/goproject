package gochan

import (
	"fmt"
	"time"
)

func ChanInit() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan=%v Chan本身的地址=%p\n", intChan, &intChan)

	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50

	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	value1 := <-intChan
	fmt.Printf("从channel取出的值=%v\n", value1)
	value2 := <-intChan
	fmt.Printf("从channel取出的值=%v\n", value2)
	value3 := <-intChan
	fmt.Printf("从channel取出的值=%v\n", value3)

	value4 := <-intChan
	fmt.Printf("从channel取出的值=%v\n", value4)
}

func ChanDemo01() {
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	close(intChan2)
	for v := range intChan2 {
		fmt.Println(v)
	}
}

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("WriteData=", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("ReadData=", v)
	}
	exitChan <- true
	close(exitChan)
}

func ChanDemo02() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool

	for {
		time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i <= num/2; i++ {
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

func GoChanPrime() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	go putNum(intChan)

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	i := 0
	for v := range primeChan {
		i++
		fmt.Println("素数=", v)
	}
	fmt.Println("素数个数=", i)

}

func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

func recv(ch <-chan int, exitChan chan struct{}) {
	for v := range ch {
		fmt.Println("接收到的数据=", v)
	}
	var a struct{}
	exitChan <- a
}

func ChanDemo03() {
	var ch chan int = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total += 1
		if total == 2 {
			break
		}
	}
	fmt.Println("over")

}

func ChanDemo04() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- fmt.Sprintf("str=%d", i)
	}

	for {
		select {
		case v := <-intChan:
			fmt.Println("从intChan读取的数据=", v)
			time.Sleep(time.Millisecond * 10)
		case v := <-stringChan:
			fmt.Println("从stringChan读取的数据=", v)
			time.Sleep(time.Millisecond * 10)
		default:
			fmt.Println("都没有数据可读，退出")
			return
		}
	}
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello Goroutine!")
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到异常:", err)
		}
	}()

	var myMap map[int]string
	myMap[1] = "one" // 这里会引发panic，因为myMap没有初始化
}

func GoChanRecover() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}
