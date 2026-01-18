package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"user-api/internal/common"
	"user-api/internal/model"
	"user-api/internal/service"
)

func LoginHandler(ctx *gin.Context) {
	var requestBody common.LoginRequest
	if err := ctx.ShouldBind(&requestBody); err != nil {
		ctx.JSON(http.StatusNotAcceptable, gin.H{"Message": "Invalid request body"})
		return
	}
	res, err := service.Login(requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.Set("user_id", requestBody.UserName)
	ctx.JSON(200, res)
}

func GetUser(ctx *gin.Context) {

	user := &model.User{UserId: 101,
		UserName:  "tushar",
		UserEmail: "tushar@gmail.com"}
	value, _ := ctx.Get("user_id")
	fmt.Println(value)
	ctx.JSON(200, user)

}
