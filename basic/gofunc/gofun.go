package gofunc

import (
	"fmt"
	"log"
	"sort"
)

func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

// 求两个数的差
func subFn(x int, y int) int {
	sub := x - y
	return sub
}

// 函数参数的简写
func subFn1(x, y int) int {
	sub := x - y
	return sub
}

// 函数的可变参数，可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识
func sumFn1(x ...int) int {
	fmt.Printf("%v...%T", x, x)
	fmt.Println("")
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func sumFn2(x int, y ...int) int {

	fmt.Println(x, y)
	log.Fatal(x, y)
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func sumFn3(x, y int) int {
	return x + y
}

// return 关键词一次可以返回多个值

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名：函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过 return 关键字 返回
func calc1(x, y int) (sum int, sub int) {
	log.Println("-------------")
	log.Println(x, y)
	sum = x + y
	sub = x - y
	return
}

func calc2(x, y int) (sum, sub int) {
	log.Println("-------------")
	log.Println(x, y)
	sum = x + y
	sub = x - y
	return
}

// int类型升序排序
func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

// int类型降序排序
func sortIntDesc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}
func mapSort(map1 map[string]string) string {
	var sliceKey []string
	for k, _ := range map1 {
		sliceKey = append(sliceKey, k)
	}
	// 对 sliceKey进行升序排序
	sort.Strings(sliceKey)
	var result string
	for _, k := range sliceKey {
		result += k + "=" + map1[k] + "&"
	}
	return result
}
