package main

/*
1.go mod init 项目名称   初始化项目
2.配置第三方包
3.go mod tidy 下载依赖
4.go run mainold.go
*/
import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

func main() {
	fmt.Println("你好 go")

	price, err := decimal.NewFromString("1.2311")
	if err != nil {
		fmt.Println("decimal.NewFromString failed")
		return
	}
	fmt.Println(price)
	var str = `{"id":1,"gender":"男","name":"张三","sno":"S0001"}`
	value := gjson.Get(str, "name")
	fmt.Println(value)
}
