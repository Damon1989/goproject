package main

// golang中结构体实例是独立的，不会影响其他实例
import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%s 年龄：%d 性别：%s \n", p.Name, p.Age, p.Sex)
}

func (p *Person) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

func main() {
	var p1 = Person{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}

	p1.PrintInfo()
	p1.SetInfo("李四", 20)

	p1.PrintInfo()

	var p2 = Person{
		Name: "王五",
		Age:  22,
		Sex:  "女",
	}

	p2.PrintInfo()

}
