package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Millisecond)
	fmt.Println(time.Second)
	fmt.Println(time.Minute)

	timeObj := time.Now()
	add := timeObj.Add(time.Hour)
	fmt.Println(timeObj)
	fmt.Println(add)
}
