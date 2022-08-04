package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
 )

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.StaticFS("/files", http.Dir("./views"))

	r.LoadHTMLGlob("views/*")
	r.GET("/", func(c *gin.Context) {
	   c.HTML(http.StatusOK, "index.html", gin.H{"content": "hello world!"})
	})
	
	t := r.Group("download")
	{
		t.GET("file", HandleDownloadFile)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}


func HandleDownloadFile(c *gin.Context) {
	content := c.Query("content")

	content = "hello world, 我是一个文件，" + content

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename=hello.txt")
	c.Header("Content-Type", "application/text/plain")
	c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	c.Writer.Write([]byte(content))
}