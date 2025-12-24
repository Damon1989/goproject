package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println("时间戳转日期：", timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Println(str1 string, str2 string) string {
	fmt.Println(str1 + str2)
	return str1 + "-----------" + str2
}

func main() {
	r := gin.Default()

	// 自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})

	//加载模板
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title":   "首页",
			"content": "哈哈哈",
			"score":   89,
			"hobby":   []string{"篮球", "足球", "乒乓球"},
			"newList": []Article{
				{
					Title:   "新闻标题1",
					Content: "新闻内容1",
				},
				{
					Title:   "新闻标题2",
					Content: "新闻内容2",
				},
			},
			"testSlice": []string{},
			"news": Article{
				Title:   "新闻标题33",
				Content: "新闻内容33",
			},
			"date": 1741054068,
		})
	})

	r.GET("/news", func(c *gin.Context) {
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻页面!!!",
			"news":  news,
		})
	})

	// 后台
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "首页",
		})
	})

	r.GET("/admin/news", func(c *gin.Context) {
		news := &Article{
			Title:   "admin新闻标题",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "admin新闻页面!!!",
			"news":  news,
		})
	})

	r.Run()
}
