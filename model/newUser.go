package model

import "github.com/pedropezzuol/api_test/src/controller/request"

func NewUser(userRequest request.UserRequest) User {
	return User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}
