package goembed

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
)

//go:embed files/robots.txt
var robots string

// 二进制文件

//go:embed files/logo.png
var logo []byte

func EmbedFile() {
	fmt.Println(robots)
	fmt.Println(logo)
}

// 嵌入目录
//
//go:embed files
var files embed.FS

func EmbedDir() {

	//获取目录下的文件列表
	entries, _ := files.ReadDir("files")
	for _, entry := range entries {
		info, _ := entry.Info()
		fmt.Println(entry.Name(), entry.IsDir(), info.Size())
	}
	fmt.Println("--------------------------------")
	data, _ := files.ReadFile("files/robots.txt")
	fmt.Println(string(data))

}

//go:embed static
var static embed.FS

func StaticEmbedServer() {
	// 获取嵌入的static子目录做为文件系统
	staticFs, _ := fs.Sub(static, "static")
	// 基于static的FS 创建http.FS
	// 基于http.FS 创建http.FileServer
	// 启动HTTP服务
	http.ListenAndServe(":8080", http.FileServer(http.FS(staticFs)))
}

// 非嵌入，运行时读取文件
func StaticRuntimeServer() {
	// os.DirFS 获取指定目录做为文件系统
	staticFs := os.DirFS("static")
	// 基于http.FS 创建http.FileServer
	// 启动HTTP服务
	http.ListenAndServe(":8081", http.FileServer(http.FS(staticFs)))
}
