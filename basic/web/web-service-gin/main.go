package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"web-service-gin/pdf2image"

	"github.com/gin-gonic/gin"
)

type album struct {
	IID    string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{IID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{IID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{IID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.IID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	fmt.Println("starting server...")
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.GET("/hello", sayHello)

	router.GET("/book", getBook)
	router.POST("/book", postBook)
	router.PUT("/book", putBook)
	router.DELETE("/book", deleteBook)

	router.GET("/pdf2image", pdf2imageProcess)

	fmt.Println("server started")
	router.Run(":8090")

}

func pdf2imageProcess(c *gin.Context) {
	pdf2image.PDFToImage("./1.pdf", "./output", 300.0, 1.0)
	c.JSON(http.StatusOK, gin.H{"method": "pdf2image"})
}

func deleteBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "delete book"})

}

func putBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "put book"})

}

func postBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "post book"})

}

func getBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "get book"})
}

func sayHello(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(c.Writer, "damon")
}
