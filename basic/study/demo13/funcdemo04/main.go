package main

import (
	"fmt"
	"sort"
)

func mapSort(map1 map[string]string) string {
	var sliceKey []string
	for key := range map1 {
		sliceKey = append(sliceKey, key)
	}
	fmt.Println(sliceKey)
	sort.Strings(sliceKey)
	//for i := 0; i < len(sliceKey); i++ {
	//	for j := i + 1; j < len(sliceKey); j++ {
	//		if sliceKey[i] > sliceKey[j] {
	//			sliceKey[i], sliceKey[j] = sliceKey[j], sliceKey[i]
	//		}
	//	}
	//}
	fmt.Println(sliceKey)
	var str string
	for _, v := range sliceKey {
		fmt.Println(v, map1[v])
		str += fmt.Sprintf("%s=>%s", v, map1[v])
	}
	fmt.Println(str)
	return str
}

func main() {
	m1 := map[string]string{
		"username": "zhangsan",
		"age":      "20",
		"sex":      "男",
		"height":   "180",
	}
	str := mapSort(m1)
	fmt.Println(str)
}
