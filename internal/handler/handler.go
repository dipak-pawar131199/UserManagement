package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"user-api/internal/common"
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
	ctx.JSON(200, res)
}
