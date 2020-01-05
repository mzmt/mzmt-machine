package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/list", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "午後の紅茶",
			"price": 120,
			"kind": "hot",
		})
	})
	r.Run()
}
