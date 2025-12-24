package main

import (
	"fmt"
	"time"
)

func main() {
	//ticker := time.NewTicker(1 * time.Second)
	//for i := range ticker.C {
	//	fmt.Println(i)
	//}

	/*	ticker := time.NewTicker(1 * time.Second)
		n := 5
		for i := range ticker.C {
			fmt.Println(i)
			n--
			if n == 0 {
				ticker.Stop()
				break
			}
		}*/

	/*	fmt.Println("aaa")
		time.Sleep(time.Second)
		fmt.Println("aaa")
		time.Sleep(time.Second)
		fmt.Println("aaa")
		time.Sleep(time.Second)
		fmt.Println("aaa")
		time.Sleep(time.Second)
		fmt.Println("aaa")
		time.Sleep(time.Second)
		fmt.Println("aaa")*/

	for {
		time.Sleep(time.Second)
		fmt.Println("aaaa")
	}
}
