package main

import "fmt"

// int 类型升序排序
func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

// int 类型降序排序
func sortIntDesc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

func main() {
	var sliceA = []int{12, 34, 37, 556, 36, 2}
	arr := sortIntAsc(sliceA)
	fmt.Println(arr)

	var sliceB = []int{1, 34, 4, 35, 6, 36, 2}
	arr1 := sortIntAsc(sliceB)
	fmt.Println(arr1)

	var sliceC = []int{1, 34, 4, 35, 6, 36, 2}
	arr2 := sortIntDesc(sliceC)
	fmt.Println(arr2)
}
