package main

import (
	"fmt"
	"sync"
)

func main() {
	defer func() {
		fmt.Println("Try to recover the panic.")
		if r := recover(); r != nil {
			fmt.Println("Recovered the panic(%#v).\n", r)
		}
	}()
	var mutex sync.Mutex
	fmt.Println("Lock the lock.(main goroutine)")
	mutex.Lock()
	fmt.Println("The lock is locked.(main goroutine)")
	fmt.Println("Unlock the lock.(main goroutine)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked.(main goroutine)")
	fmt.Println("Unlock the lock again" +
		"" +
		".(main goroutine)")
}
