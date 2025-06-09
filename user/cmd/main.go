package main

import (
	"user/internal/handler"
	"user/internal/logic"
)

func main() {
	//config.InitConfig()
	//repository.InitDB()
	service := handler.UserService{}
	service.UserLogin(nil, &logic.UserRequest{
		NickName:        "lc6",
		UserName:        "lc",
		Password:        "123456",
		PasswordConfirm: "123456",
	})
}
