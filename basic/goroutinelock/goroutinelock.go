package goroutinelock

// go build -race -o goroutinelock.exe goroutinelock.go goroutinelock_test.go
import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup
var mutex sync.Mutex
var rwMutex sync.RWMutex

func test() {
	mutex.Lock()
	defer mutex.Unlock()
	count++
	fmt.Println("the count is :", count)
	time.Sleep(time.Millisecond)
	wg.Done()
}

func lockMain() {
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}

func write() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	fmt.Println("write")
	time.Sleep(time.Second * 2)
	wg.Done()
}

func read() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	fmt.Println("read")
	time.Sleep(time.Second * 2)
	wg.Done()
}

func readWriteMain() {
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
