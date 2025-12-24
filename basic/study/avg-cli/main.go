package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入一系列数字（以空格分隔）：")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// 将输入字符串分割成数字切片
	numbers := strings.Split(input, " ")

	var nums []float64

	for _, numStr := range numbers {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			fmt.Println("无法解析输入的数字：", numStr)
			return
		}
		nums = append(nums, num)
	}

	sum := 0.0
	for _, num := range nums {
		sum += num
	}

	avg := sum / float64(len(nums))

	fmt.Printf("这些数字的平均值是：%0.2f\n", avg)
}
