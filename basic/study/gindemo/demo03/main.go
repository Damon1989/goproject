package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()

	//加载模板
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"score": 89,
			"hobby": []string{"篮球", "足球", "乒乓球"},
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
