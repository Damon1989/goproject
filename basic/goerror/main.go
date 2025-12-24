package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	//f1()
	//f2()
	//f3()

	//f4()
	//f5()

	half, err := Half(18)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Half is:", half)
	}
	fmt.Println(half)
}

func f5() {
	half, err := Half(19)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Half is:", half)
	}
	fmt.Println(half)
}

func Half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, errors.New("Cannot half an odd number")
	}
	return numberToHalf / 2, nil
}

func f4() {
	name, role := "lnj", "admin"
	err := fmt.Errorf("The user %s has role %s", name, role)
	if err != nil {
		fmt.Println(err)
	}
}

func f3() {
	err := errors.New("This is an error!!!.")
	if err != nil {
		fmt.Println(err)
	}
}

func f2() {
	file, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(file))
}

func f1() {
	err := fmt.Errorf("This is an error.")
	fmt.Println(err)

	err = errors.New("This is an error！！！.")
	fmt.Println(err)
}

//	func div(a, b int) (int, error) {
//		if b == 0 {
//			return 0, errors.New("Division by zero.")
//		}
//		return a / b, nil
//	}
//func div(a, b int) (res int) {
//	if b == 0 {
//		panic("Division by zero.")
//	}
//	return a / b
//}

func main2() {
	////res, err := div(10, 0)
	//res, err := div(10, 5)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(res)
	//}
	res := div(10, 0)
	fmt.Println(res)
}

func main3() {
	var arr = [3]int{1, 2, 3}
	//arr[5] = 5
	fmt.Println(arr)

	//var res = 10 / 0
	//fmt.Println(res)
}

func div(a, b int) (res int) {
	// 定义一个延迟调用的函数，用于捕获panic异常
	// 注意：recover()函数必须在defer语句中调用才有效
	// 一定要在panic之前定义
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			res = -1
		}
	}()
	if b == 0 {
		panic("Division by zero.")
	}
	return a / b
}

func test2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // This is B panic.
		}
	}()
	defer func() {
		panic("This is B panic.")
	}()
	panic("This is A panic.")
}

func test3() {
	err := fmt.Errorf("这里是错误信息")
	fmt.Println(err)
}

func test4() {
	err := errors.New("这里是错误信息!!!")
	fmt.Println(err)
}

func myDiv(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("Division by zero.")
	} else {
		res = a / b
	}
	return
}

func myDiv1(a, b int) (res int) {
	if b == 0 {
		// 一旦传入的除数为0，程序就会终止
		panic("Division by zero.")
	} else {
		res = a / b
	}
	return
}

func myDivRecover(a, b int) (res int) {
	// 定义一个延迟调用的函数，用于捕获panic异常
	// 注意：一定要在panic之前定义
	defer func() {
		if err := recover(); err != nil {
			res = -1
			fmt.Println(err) // 除数不能为0
		}
	}()
	if b == 0 {
		panic("Division by zero.")
	} else {
		res = a / b
	}
	return
}

func myRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		panic("This is B panic.")
	}()
	panic("This is A panic.")
}

func f() {
	//res := div(10, 0)
	//fmt.Println(res)
	//test2()
	//test3()
	//test4()
	/*	res, err := myDiv(10, 2)
		//res, err := myDiv(10, 0)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}*/

	/*	res := myDiv1(10, 0)
		fmt.Println(res)*/
	//myDivRecover(10, 0)
	myRecover()
}
