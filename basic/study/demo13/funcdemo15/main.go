package main

import (
	"fmt"
	"strings"
	"time"
)

func f1() {
	fmt.Println("fn1")
}

func f2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("恢复异常：", err)
		}
	}()
	panic("抛出一个异常")
}

func main() {
	//f1()
	//f2()
	//fmt.Println("结束")
	//err := fmt.Errorf("这里是错误信息")
	//fmt.Println(err)
	//
	//err = errors.New("这里是错误信息！")
	//fmt.Println(err)
	//
	//println("结束")
	//
	//str := "公号：代码情缘"
	//fmt.Println(str)
	//arr1 := []byte(str)
	////fmt.Println(arr1)
	//
	//for _, v := range arr1 {
	//	fmt.Printf("%c", v)
	//}
	//
	//arr2 := []rune(str)
	//fmt.Println(arr2)
	//for _, v := range arr2 {
	//	fmt.Printf("%c", v)
	//}

	res := strings.Compare("bcd", "abc")
	fmt.Println(res)

	now := time.Now()
	fmt.Println(now)
}
