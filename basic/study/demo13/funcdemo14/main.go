package main

/*
*
提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
*/
func calc(index string, a, b int) int {
	ret := a + b
	println(index, a, b, ret)
	return ret
}

func main() {

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
