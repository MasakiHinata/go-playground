package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	// ミドルウェア
	r.Use(ObserveUserAgent)

	// ルーティング
	r.GET("/", HelloWorld)
	r.POST("/hello", Hello)
	r.GET("/cookie", Cookie)
	r.GET("/secure_json", SecureJSON)

	v2 := r.Group("/v2")
	{
		v2.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"version": 2,
			})
		})
	}

	r.Run(":8080")
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

type Person struct {
	// 必須パラメータには binding"required" を付与
	Name string `json:"name" binding:"required"`
}

func Hello(c *gin.Context) {
	req := Person{}
	if err := c.Bind(&req); err != nil {
		http.Error(c.Writer, fmt.Sprintf("error while parsing request: %v", err), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Hello %s!", req.Name))
}

func ObserveUserAgent(c *gin.Context) {
	log.Println("User-Agent:", c.GetHeader("User-Agent"))
	c.Next()
}

func Cookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}

	c.JSON(http.StatusOK, gin.H{
		"cookie value": cookie,
	})
}

func SecureJSON(c *gin.Context) {
	names := []string{"lena", "austin", "foo"}

	// Will output  :   while(1);["lena","austin","foo"]
	c.SecureJSON(http.StatusOK, names)
}
