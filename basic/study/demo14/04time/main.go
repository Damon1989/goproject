package main

import (
	"fmt"
	"time"
)

func main() {
	var str = "2025-02-20 18:05:04"
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	fmt.Println(timeObj, err)

	fmt.Println(timeObj.Unix())
}
