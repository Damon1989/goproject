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

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d, Score: %d", s.Name, s.Age, s.Score)
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) PrintInfo() {
	fmt.Println("这是一个学生信息打印方法")
}

func printStructField(s interface{}) {
	t := reflect.TypeOf(s)
	// Check if it's a struct or a pointer to struct
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("Expected a struct type")
		return
	}
	field0 := t.Field(0)
	fmt.Printf("Field Name: %s, Field Type: %s, Tag: %s ,TagJsonName:%s \n", field0.Name, field0.Type, field0.Tag,
		field0.Tag.Get("json"))
	fmt.Println(field0)
	fmt.Printf("%#v\n", field0)

	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Printf("Field1 Name: %s, Field1 Type: %s, Tag1: %s \n", field1.Name, field1.Type, field1.Tag)
	}

	fmt.Println("----遍历结构体的所有字段----")

	var fieldCount = t.NumField()
	fmt.Printf("Total number of fields: %d\n", fieldCount)

	for i := 0; i < fieldCount; i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: Name: %s, Type: %s, Tag: %s \n", i, field.Name, field.Type, field.Tag)
	}
	fmt.Println("遍历获取结构体字段值")
	v := reflect.ValueOf(s)
	for i := 0; i < fieldCount; i++ {
		fieldValue := v.Field(i)
		fmt.Printf("Field %d: Value: %v \n", i, fieldValue.Interface())
	}

}

func printStructMethod(s interface{}) {
	t := reflect.TypeOf(s)
	// Check if it's a struct or a pointer to struct
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("Expected a struct type")
		return
	}
	fmt.Println("----遍历结构体的方法----")
	methodCount := t.NumMethod()
	fmt.Printf("Total number of methods: %d\n", methodCount)
	for i := 0; i < methodCount; i++ {
		method := t.Method(i)
		fmt.Printf("Method %d: Name: %s, Type: %s \n", i, method.Name, method.Type)
	}

	v := reflect.ValueOf(s)
	v.Method(1).Call(nil)
	fmt.Println(v.Method(0).Call(nil))

	v.MethodByName("SetInfo").Call([]reflect.Value{reflect.ValueOf("Charlie"), reflect.ValueOf(22), reflect.ValueOf(95)})
	fmt.Println(v.MethodByName("GetInfo").Call(nil))
}

func reflectChangeStructField(s interface{}) {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		fmt.Println("Expected a pointer to struct type")
		return
	}
	v := reflect.ValueOf(s).Elem()
	field := v.FieldByName("Name")
	if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString("David")
	}
	ageField := v.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() && ageField.Kind() == reflect.Int {
		ageField.SetInt(25)
	}
	scoreField := v.FieldByName("Score")
	if scoreField.IsValid() && scoreField.CanSet() && scoreField.Kind() == reflect.Int {
		scoreField.SetInt(88)
	}
	fmt.Printf("Modified struct: %+v\n", s)

}

func reflectStructMain() {
	student := Student{Name: "Bob", Age: 20, Score: 90}
	//printStructField(student)
	//printStructMethod(&student)
	reflectChangeStructField(&student)
}
