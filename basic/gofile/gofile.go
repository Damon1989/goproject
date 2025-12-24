package gofile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func openFile() {
	// 只读方式
	file, err := os.Open("./test.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("文件打开成功:", file.Name())

	var tempSlice = make([]byte, 128)
	n, err1 := file.Read(tempSlice)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("读取了多少字节:", n)
	fmt.Println("读取的内容是:", string(tempSlice))
	// 另一种读取文件内容的方式
	data, err := os.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("读取的内容是:", string(data))
}

func OpenFile1() {
	file, err := os.Open("./test.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("文件打开成功:", file.Name())

	var fileStr string
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fileStr += line
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fileStr += line
	}
	fmt.Println(fileStr)
}

func OpenFile2() {
	bytes, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func WriteFile() {
	file, err := os.OpenFile("./test_write.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("文件打开成功:", file.Name())

	content := "Hello, Go File Writing!\nThis is a test file."

	n, err := file.Write([]byte(content))
	if err != nil {
		panic(err)
	}
	fmt.Printf("写入了 %d 个字节\n", n)
}

func WriteFile1() {
	file, err := os.OpenFile("./test_write1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("文件打开成功:", file.Name())

	writer := bufio.NewWriter(file)
	content := "Hello, Go Buffered Writing!\nThis is another test file."

	n, err := writer.WriteString(content)
	if err != nil {
		panic(err)
	}
	writer.Flush()
	fmt.Printf("写入了 %d 个字节\n", n)

}

func WriteFile2() {
	content := "Hello, Go ioutil Writing!\nThis is yet another test file."
	err := ioutil.WriteFile("./test_write2.txt", []byte(content), 0666)
	if err != nil {
		panic(err)
	}
	fmt.Println("文件写入成功")
}
