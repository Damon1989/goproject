package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前计算机的CPU数量
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum:", cpuNum)

	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
