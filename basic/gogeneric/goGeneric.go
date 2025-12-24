package goGeneric

import (
	"fmt"
	"sync"
)

func GenericDef() {
	// 一：泛型类型的定义
	//	使用类型参数来定义类型
	type mySlice[P int | string] []P
	// 泛型类型 mySlice[P]

	type myMap[K int | string, V float64 | float32] map[K]V

	type myList[T int | float64 | string] struct {
		data          []T
		l             int
		max, min, avg T
	}

	// 二：声明泛型类型的变量
	// 使用具体类型来实例化泛型类型
	_ = mySlice[int]{1, 2, 3, 4, 5}
	_ = mySlice[string]{"a", "b", "c"}
	_ = myMap[string, float64]{"a": 1.1, "b": 2.2}
	_ = myList[float64]{data: []float64{1.1, 2.2, 3.3}, l: 3, max: 3.3, min: 1.1, avg: 2.2}

	fmt.Println("GenericDef executed")
}

func GenericExt() {
	// 一：类型约束为基础类型
	// ~

	type intSlice[P ~int] []P
	_ = intSlice[int]{}

	// 以int为基础类型的自定义类型
	type myInt int
	_ = intSlice[myInt]{}
	type yourInt int
	_ = intSlice[yourInt]{}

	type mySlice[T int | string | float64 | float32] []T
	_ = mySlice[int]{1, 2, 3}
	_ = mySlice[string]{"a", "b", "c"}
	_ = mySlice[float64]{1.1, 2.2, 3.3}
	_ = mySlice[float32]{1.1, 2.2, 3.3}

	type floatSlice[T float32 | float64] mySlice[T]

	type myStruct[T float32 | float64] struct {
		data mySlice[T]
		l    int
		max  T
		min  T
		avg  T
	}

	fmt.Println("GenericExt executed")

}

// 一：定义泛型类型
type myList[T int | float64] struct {
	data     []T
	max, min T
	m        sync.Mutex
}

// 二：定义泛型类型的方法集
// 1.添加元素，并更新最大值和最小值
func (l *myList[T]) Add(ele T) *myList[T] {
	l.m.Lock()
	defer l.m.Unlock()
	l.data = append(l.data, ele)
	if len(l.data) == 1 {
		l.max = ele
		l.min = ele
	} else {
		if ele > l.max {
			l.max = ele
		}
		if ele < l.min {
			l.min = ele
		}
	}
	return l
}

// 2.获取元素
func (l *myList[T]) All() []T {
	return l.data
}

func (l *myList[T]) Max() T {
	return l.max
}

func (l *myList[T]) Min() T {
	return l.min
}

func GenericReceiver() {
	list := myList[int]{}
	list.Add(10).Add(20).Add(5).Add(30)
	fmt.Println("List elements:", list.All())
	fmt.Println("Max:", list.Max())
	fmt.Println("Min:", list.Min())

	floatList := myList[float64]{}
	floatList.Add(1.5).Add(2.3).Add(0.8).Add(4.1)
	fmt.Println("Float List elements:", floatList.All())
	fmt.Println("Max:", floatList.Max())
	fmt.Println("Min:", floatList.Min())
}

func Sum[T int | string](ele ...T) T {
	var s T
	for _, v := range ele {
		s += v
	}
	return s
}

func GenericFunc() {
	fmt.Println(Sum[int](1, 2, 3, 4, 5))
	fmt.Println(Sum[string]("a", "b", "c"))
}

func GuessType[T int | string](ele ...T) T {
	var s T
	for _, v := range ele {
		s += v
	}
	return s
}

func TypeInference() {
	fmt.Println(GuessType[int](1, 2, 3))
	fmt.Println(GuessType[string]("a", "b", "c"))

	GuessType2[int, string](42, "Ma")
	GuessType2[int](42, "Hello")
	GuessType2(44, "world")
}

func GuessType2[K int | string, V float64 | string](p1 K, p2 V) {
	fmt.Printf("Type of p1: %T, value: %v\n", p1, p1)
	fmt.Printf("Type of p2: %T, value: %v\n", p2, p2)
}

type Data[T int | string] interface {
	Process(T) (T, error)
	Save() error
}

func DataOperate(p Data[string]) {
	
}
