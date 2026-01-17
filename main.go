package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"user-api/internal/config"
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
	r.Run(":8080")
}
