package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock.(main goroutine)")
	mutex.Lock()
	fmt.Println("The lock is locked.(main goroutine)")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock.(g%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked.(g%d)\n", i)
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Unlock the lock.(main goroutine)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked.(main goroutine)")
	time.Sleep(10 * time.Second)
}
