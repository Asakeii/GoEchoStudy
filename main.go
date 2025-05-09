package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// 中间件
func midStatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		log.Println(cost)
	}

}
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Use(midStatCost())

	admin := r.Group("/admin")
	{

		admin.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "demo.html", gin.H{})
		})
		admin.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"word": c.PostForm("f1"),
			})
		})
	}

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
