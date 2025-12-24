package gofunc

import (
	"fmt"
	"log"
	"testing"
)

func TestSumFn(t *testing.T) {
	sum := sumFn(1, 2)
	fmt.Println(sum)
}

func TestSubFn(t *testing.T) {
	sub := subFn(10, 2)
	fmt.Println(sub)
}

func TestSubFun1(t *testing.T) {
	sub := subFn1(10, 2)
	fmt.Println(sub)
}

func TestSumFun1(t *testing.T) {
	sum := sumFn1(1, 2, 3, 4, 5)
	fmt.Println(sum)
}

func TestSubFun2(t *testing.T) {
	sum := sumFn2(10, 2, 3, 5)
	fmt.Println(sum)
}

func TestCalc(t *testing.T) {
	sum, sub := calc(10, 2)
	log.Println(sum, sub)
}

func TestCalc1(t *testing.T) {
	sum, sub := calc1(110, 2)
	log.Println(sum, sub)
}

func TestCalc2(t *testing.T) {
	sum, sub := calc2(110, 2)
	log.Println(sum, sub)
	sum, _ = calc2(110, 2)
	log.Println(sum)
}

func TestSortIntAsc(t *testing.T) {
	arr := sortIntAsc([]int{1, 2, 3, 4, 5, 2, 3, 4})
	log.Println(arr)
}

func TestSortIntDesc(t *testing.T) {
	arr := sortIntDesc([]int{1, 2, 3, 4, 5, 2, 3, 4})
	log.Println(arr)
}
func TestMapSort(t *testing.T) {
	m1 := map[string]string{
		"username": "zhangsan",
		"age":      "20",
		"sex":      "男",
		"height":   "180cm",
	}
	log.Println(m1["username"])
	log.Println(mapSort(m1))
}
