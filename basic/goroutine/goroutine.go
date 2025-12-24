package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func GoroutineAnts() {
	// 1.统计当前存在的goroutine数量
	go func() {
		for {
			fmt.Println("NumGoroutine:", runtime.NumGoroutine())
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 2.初始化协程池，
	size := 1024
	pool, err := ants.NewPool(size)
	if err != nil {
		fmt.Println("创建协程池失败:", err)
		return
	}
	// 保证pool在函数退出时释放资源
	defer pool.Release()

	for {
		// 3.利用pool,调度需要并发的大量goroutine任务
		err = pool.Submit(func() {
			v := make([]byte, 1024)
			_ = v
			fmt.Println("in goroutine")
			time.Sleep(100 * time.Second)
		})
		if err != nil {
			fmt.Println("提交任务失败:", err)
			return
		}
	}

}

func GoroutineScheduler() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 设置为1个P在调度G
	runtime.GOMAXPROCS(1)
	max := 100
	go func() {
		defer wg.Done()
		for i := 1; i <= max; i += 2 {
			fmt.Print(i, " ")
			runtime.Gosched()
			//time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= max; i += 2 {
			fmt.Print(i, " ")
			runtime.Gosched()
			//time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
}

func slowFun() {
	time.Sleep(2 * time.Second)
	fmt.Println("sleeper() done")
}

func PrimeNumberFor() {
	start := time.Now().UnixMilli()
	for num := 2; num < 120000; num++ {
		var flag = true
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf("%d \n", num)
		}
	}
	end := time.Now().UnixMilli()
	fmt.Println("cost", end-start)
}

func PrimeNumberGoroutine() {
	start := time.Now().UnixMilli()
	var wg sync.WaitGroup
	for num := 2; num < 120000; num++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			var flag = true
			for i := 2; i <= n/2; i++ {
				if n%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				fmt.Printf("%d \n", n)
			}
		}(num)
	}
	wg.Wait()
	end := time.Now().UnixMilli()
	fmt.Println("cost", end-start)
}
