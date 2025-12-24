package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup
var mutex sync.Mutex

var m = make(map[int]int, 40)

/*func test() {
	mutex.Lock()
	count++
	fmt.Println("the count is :", count)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}*/

func test(num int) {
	mutex.Lock()
	var sum = 1
	for i := 1; i <= num; i++ {
		sum *= num
	}
	m[num] = sum

	fmt.Println(num, "的阶乘是：", sum)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
}
