package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()

	//注意：配置模板的文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "首页")
	})

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好gin",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好gin",
		})
	})

	r.GET("/json3", func(c *gin.Context) {
		a := Article{
			Title:   "标题",
			Desc:    "描述",
			Content: "内容",
		}
		c.JSON(200, a)
	})

	// 响应jsonp请求
	// /jsonp?callback=xxx
	// xxx({"title":"标题-jsonp","desc":"描述","content":"内容"});
	r.GET("/jsonp", func(c *gin.Context) {
		a := Article{
			Title:   "标题-jsonp",
			Desc:    "描述",
			Content: "内容",
		}
		c.JSONP(200, a)
	})

	r.GET("/xml", func(c *gin.Context) {
		a := Article{
			Title:   "标题-xml",
			Desc:    "描述",
			Content: "内容",
		}
		c.XML(200, a)
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台的数据",
		})
	})

	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是商品页面",
			"price": 20,
		})
	})

	r.Run()
}
