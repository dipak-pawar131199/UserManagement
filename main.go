package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"user-api/internal/config"
	"user-api/internal/handler"
)

func main() {
	fmt.Println("hello")

	d := config.ConnectionDB().DB
	fmt.Println(d)

	// create gin condex

	r := gin.Default()
	r.GET("user-api/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "hello from gin",
		})
	})
	// r.POST("user-api/login", func(ctx *gin.Context) {
	// 	service.Login(ctx)
	// })
	r.POST("user/api/login", handler.LoginHandler)

	r.Run(":8080")
}
