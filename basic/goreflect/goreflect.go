package goreflect

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	Name string
	Age  int
}

// 反射获取任意变量的类型和值
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	name := v.Name()
	kind := v.Kind()
	fmt.Printf("Type:%v,Name:%v,Kind:%v \n", v.String(), name, kind)
}

func reflectMain() {
	a := 10
	b := 23.4
	c := true
	d := "hello"

	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e myInt = 34
	reflectFn(e)

	p := Person{Name: "Alice", Age: 30}
	reflectFn(p)

	var h = 25
	reflectFn(&h)

	var i = [3]int{1, 2, 3}
	reflectFn(i)

	var j = []string{"a", "b", "c"}
	reflectFn(j)

	var k = map[string]int{"one": 1, "two": 2}
	reflectFn(k)

	var l = func(x int) int { return x * x }
	reflectFn(l)
}

func reflectValue(x interface{}) {
	//v := reflect.ValueOf(x)
	//fmt.Printf("Value:%v,Kind:%v \n", v.Interface(), v.Kind())
	/*	b, ok := x.(int)
		if ok {
			b += 10
			x = b
		}
		fmt.Println(x)*/
	/*switch v := x.(type) {
	case int:
		v += 10
		fmt.Println(v)
	case string:
		v += " world"
		fmt.Println(v)
	case float64:
		v *= 2
		fmt.Println(v)
	case bool:
		v = !v
		fmt.Println(v)
	default:
		fmt.Println("unsupported type")
	}*/
	switch v := reflect.ValueOf(x); v.Kind() {
	case reflect.Int:
		newValue := v.Int() + 10
		fmt.Println(newValue)
	case reflect.String:
		newValue := v.String() + " world"
		fmt.Println(newValue)
	case reflect.Float64:
		newValue := v.Float() * 2
		fmt.Println(newValue)
	case reflect.Bool:
		newValue := !v.Bool()
		fmt.Println(newValue)
	default:
		fmt.Println("unsupported type")
	}
}

func reflectValueMain() {
	a := 42
	b := "gopher"
	c := 3.14
	d := false

	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
	reflectValue(d)
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Cannot set value")
		return
	}
	v = v.Elem()
	switch v.Kind() {
	case reflect.Int:
		v.SetInt(100)
	case reflect.String:
		v.SetString("modified")
	case reflect.Float64:
		v.SetFloat(6.28)
	case reflect.Bool:
		v.SetBool(true)
	default:
		fmt.Println("unsupported type")
	}
}

func reflectSetValueMain() {
	a := 10
	b := "hello"
	c := 3.14
	d := false

	reflectSetValue(&a)
	reflectSetValue(&b)
	reflectSetValue(&c)
	reflectSetValue(&d)

	fmt.Println("Modified values:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
}
