package gofiledirectory

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CopyFile(srcFile string, dstFile string) {
	byteStr, err1 := ioutil.ReadFile(srcFile)
	if err1 != nil {
		panic(err1)
	}
	err2 := ioutil.WriteFile(dstFile, byteStr, 0666)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("文件复制成功")
}

func CopyFileStream(srcFile string, dstFile string) error {
	sFile, err1 := os.Open(srcFile)
	dFile, err2 := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer sFile.Close()
	defer dFile.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	var tempSlice = make([]byte, 128)
	for {
		n, err := sFile.Read(tempSlice)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				return err
			}
		}
		_, err = dFile.Write(tempSlice[:n])
		if err != nil {
			return err
		}
	}
	return nil

}

func CreateDir(dirName string) {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		//panic(err)
		fmt.Println("目录创建失败:", err)
	}
	fmt.Println("目录创建成功:", dirName)
}

func CreateDirAll(dirName string) {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		//panic(err)
		fmt.Println("目录创建失败:", err)
	}
	fmt.Println("目录创建成功:", dirName)
}

func RemoveDir(dirName string) {
	err := os.Remove(dirName)
	if err != nil {
		//panic(err)
		fmt.Println("目录删除失败:", err)
	}
	fmt.Println("目录删除成功:", dirName)
}
func RemoveDirAll(dirName string) {
	err := os.RemoveAll(dirName)
	if err != nil {
		//panic(err)
		fmt.Println("目录删除失败:", err)
	}
	fmt.Println("目录删除成功:", dirName)
}
