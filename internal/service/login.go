package service

import (
	"errors"
	"net/http"

	"user-api/internal/common"
	"user-api/internal/config"
	"user-api/internal/model"
	"user-api/internal/utils"
)

type Response struct {
	Message string
	Status  int
	Token   string
}

func Login(requestBody common.LoginRequest) (*Response, error) {

	db := config.ConnectionDB().DB
	var user model.User

	err := db.QueryRow("select userName,userEmail from user where userName = ? and password =?", requestBody.UserName, requestBody.Password).Scan(&user.UserName, &user.UserEmail)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}

	token, _ := utils.GenerateToken(user.UserName)

	return &Response{Message: "Login Sucesses", Status: http.StatusOK, Token: token}, nil
}
