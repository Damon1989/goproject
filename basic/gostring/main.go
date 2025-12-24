package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//f1()

	var buffer bytes.Buffer
	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}
	fmt.Println(buffer.String())
}

func f1() {
	res := strings.IndexByte("hello 李南江", 'l')
	fmt.Println(res)

	res = strings.IndexRune("hello 李南江", '李')
	fmt.Println(res)

	res = strings.IndexRune("hello 李南江", 'l')
	fmt.Println(res)

	res = strings.IndexAny("hello 李南江", "wml")
	fmt.Println(res)
}

func main1() {
	str1 := "lnj"
	fmt.Println(len(str1))
	str2 := "公号:代码情缘"
	fmt.Println(len(str2))
}

func main2() {
	str := "公号:代码情缘"
	arr1 := []byte(str)
	fmt.Println(len(arr1))
	for _, v := range arr1 {
		fmt.Printf("%c", v)
	}

	arr2 := []rune(str)
	fmt.Println(len(arr2))
	for _, v := range arr2 {
		fmt.Printf("%c", v)
	}
}
