package main

import "fmt"

func main1() {
	var dic map[int]int = map[int]int{1: 1, 2: 2, 3: 3}
	fmt.Println(dic)
	fmt.Println(dic[1])

	dic[1] = 666
	fmt.Println(dic)
}

func main2() {
	dict := map[string]string{
		"Tom":   "男",
		"Jerry": "女",
	}
	fmt.Println(dict)

	var dict1 = make(map[string]string, 2)
	dict1["Tom"] = "男"
	dict1["Jerry"] = "女"
	fmt.Println(dict1)

	var dict2 = make(map[string]string)
	dict2["Tom"] = "男1"
	dict2["Jerry"] = "女1"
	fmt.Println(dict2)

	var dict3 = make(map[string]string)
	fmt.Println("增加前", dict3)
	dict3["Tom"] = "男1"
	dict3["Jerry"] = "女1"
	fmt.Println("增加后", dict3)

}

func main() {
	var dict = map[string]string{
		"Tom":   "男",
		"Jerry": "女",
		"name":  "damon",
	}
	for key, value := range dict {
		fmt.Println(key, value)
	}
}
