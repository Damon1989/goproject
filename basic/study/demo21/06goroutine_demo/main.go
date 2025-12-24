package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：要统计1-120000的数字中哪些是素数？goroutine for循环

/*
*
1 协程 统计 1-30000
2 协程 统计 30001-60000
3 协程 统计 60001-90000
4 协程 统计 90001-120000
*/
// start:(n-1)*30000+1 end:n*30000

var wg sync.WaitGroup

func test(n int) {
	defer wg.Done()
	for num := (n-1)*30000 + 1; num < n*30000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				fmt.Printf("%d是素数\n", num)
			}
		}
	}
}
func main() {
	start := time.Now().UnixMicro()
	for i := 1; i <= 4; i++ {
		go test(i)
		wg.Add(1)
	}
	wg.Wait()
	end := time.Now().UnixMicro()
	fmt.Println("耗时：", end-start)
}
