package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	db := NewConfig()
	// defer db.Close()
	// db.PingTest()
	db.Migration(&Todo{})

	rg := r.Group("/api")
	rg.GET("/todos", func(ctx *gin.Context) {
		var todo []Todo
		if err := db.Db.Find(&todo).Error; err != nil {
			ctx.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			ctx.JSON(200, gin.H{
				"message": "Ok",
				"data":    todo,
			})
		}
	})
	r.Run(":8081")
}
