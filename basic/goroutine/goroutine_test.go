package goroutine

import (
	"fmt"
	"testing"
)

func TestGoroutineAnts(t *testing.T) {
	GoroutineAnts()
}

func TestGoroutineScheduler(t *testing.T) {
	GoroutineScheduler()
}

func TestSlowFun(t *testing.T) {
	slowFun()
	fmt.Println("I am not shown until slowFun ends")
}

func TestPrimeNumberFor(t *testing.T) {
	PrimeNumberFor()
}

func TestPrimeNumberGoroutine(t *testing.T) {
	PrimeNumberGoroutine()
}
