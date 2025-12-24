package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

type UserInfo struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func main() {
	r := gin.Default()

	//加载模板
	r.LoadHTMLGlob("templates/**/*")

	// Get 请求传值
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		page := c.DefaultQuery("page", "1")
		fmt.Println("username:", username)
		fmt.Println("page:", page)
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"page":     page,
		})
	})

	r.GET("/article", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情",
			"id":  id,
		})
	})

	// post
	r.POST("/user", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "18")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	// 获取 get post 传递的数据绑定到结构体
	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
	})

	r.POST("/postUser", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
	})

	// 获取post xml 数据
	r.POST("/postXml", func(c *gin.Context) {
		xmlSliceData, _ := c.GetRawData() // 获取c.Request.Body的数据
		article := &Article{}
		fmt.Println(xmlSliceData)
		fmt.Println("xmlSliceData:", string(xmlSliceData))
		if err := xml.Unmarshal(xmlSliceData, article); err == nil {
			fmt.Printf("%#v\n", article)
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
	})

	//动态路由传值
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	r.Run()
}
